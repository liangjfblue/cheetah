package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/pkg/errors"

	"go.etcd.io/etcd/clientv3"
)

type Master struct {
	Etcd       *Etcd
	nodeSelect Select //select algorithm
	mutex      *sync.RWMutex
	allNodes   []*Member
	srvNodes   map[string][]*Member
	pathNodes  map[string][]*Member
}

func NewMaster(etcd *Etcd, nodeSelect Select) *Master {
	return &Master{
		Etcd:       etcd,
		nodeSelect: nodeSelect,
		mutex:      &sync.RWMutex{},
		allNodes:   make([]*Member, 0),
		srvNodes:   make(map[string][]*Member, 0),
		pathNodes:  make(map[string][]*Member, 0),
	}
}

//WatchEvent watch path and wait put/del event
func (m *Master) WatchEvent(ctx context.Context, path string) {
	watcher := clientv3.NewWatcher(m.Etcd.Client)
	watchChan := watcher.Watch(ctx, path, clientv3.WithFromKey(), clientv3.WithPrefix())

	for {
		select {
		case event := <-watchChan:
			for _, ev := range event.Events {
				fmt.Println(ev)
				switch ev.Type {
				case clientv3.EventTypePut:
					m.putEvent(ev)
				case clientv3.EventTypeDelete:
					m.delEvent(ev)
				}
			}
		}
	}
}

func (m *Master) putEvent(ev *clientv3.Event) {
	var (
		nodeInfo NodeInfo
	)

	if err := json.Unmarshal(ev.Kv.Value, &nodeInfo); err != nil {
		log.Fatal(err)
	}

	member := Member{
		Online:   true,
		NodeInfo: nodeInfo,
	}

	m.add(&member)
}

func (m *Master) delEvent(ev *clientv3.Event) {
	var (
		nodeInfo NodeInfo
	)

	if err := json.Unmarshal(ev.Kv.Value, &nodeInfo); err != nil {
		log.Fatal(err)
	}

	member := Member{
		Online:   true,
		NodeInfo: nodeInfo,
	}
	if err := m.del(&member); err != nil {
		log.Fatal(err.Error())
	}
}

//add add node into master allNodes map
func (m *Master) add(member *Member) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.pathNodes[member.Path] == nil {
		m.pathNodes[member.Path] = make([]*Member, 0)
	}
	m.pathNodes[member.Path] = append(m.pathNodes[member.Path], member)

	if m.srvNodes[member.SrvName] == nil {
		m.srvNodes[member.SrvName] = make([]*Member, 0)
	}
	m.pathNodes[member.SrvName] = append(m.pathNodes[member.SrvName], member)

	m.allNodes = append(m.allNodes, member)
}

//del delete node from master allNodes map
func (m *Master) del(member *Member) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if all, ok := m.srvNodes[member.SrvName]; ok {
		for k, node := range all {
			if member.Hostname == node.Hostname &&
				member.Env == node.Env &&
				member.Path == node.Path {
				m.srvNodes[member.SrvName] = append(m.srvNodes[member.SrvName][:k], m.srvNodes[member.SrvName][k+1:]...)
			}
		}
	}

	if all, ok := m.pathNodes[member.Path]; ok {
		for k, node := range all {
			if member.Hostname == node.Hostname &&
				member.Env == node.Env &&
				member.Path == node.Path {
				m.pathNodes[member.Path] = append(m.pathNodes[member.Path][:k], m.pathNodes[member.Path][k+1:]...)
			}
		}
	}

	for k, node := range m.allNodes {
		if member.Hostname == node.Hostname &&
			member.Env == node.Env &&
			member.Path == node.Path {
			m.allNodes = append(m.allNodes[:k], m.allNodes[k+1:]...)
		}
	}
	return nil
}

//SrvAll get all node from watch path
func (m *Master) SrvAll(srvName string) []*Member {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if all, ok := m.srvNodes[srvName]; ok {
		return all
	} else {
		return nil
	}

}

//get one online node by srvName
func (m *Master) Get(srvName string) (*Member, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if all, ok := m.srvNodes[srvName]; ok {
		return m.nodeSelect.Index(all)
	} else {
		return nil, errors.New("no this service nodes")
	}
}
