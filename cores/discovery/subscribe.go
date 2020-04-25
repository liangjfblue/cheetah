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
	Etcd     *Etcd
	mutex    *sync.RWMutex
	srvNodes map[string]NodeInfo
}

func NewMaster(etcd *Etcd) *Master {
	return &Master{
		Etcd:     etcd,
		mutex:    &sync.RWMutex{},
		srvNodes: make(map[string]NodeInfo, 0),
	}
}

//WatchEvent watch path and wait put/del event
func (m *Master) Watch(ctx context.Context, path string) {
	watchChan := m.Etcd.Watch(ctx, path, clientv3.WithFromKey(), clientv3.WithPrefix())

	for {
		select {
		case event := <-watchChan:
			for _, ev := range event.Events {
				log.Println(ev)
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

//All get all node from watch path
func (m *Master) All(ctx context.Context, srvName string) (NodeInfo, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if node, ok := m.srvNodes[srvName]; ok {
		return node, nil
	} else {
		return NodeInfo{}, errors.New(fmt.Sprintf("service:%s empty nodes", srvName))
	}

}

//Get one online node by srvName
func (m *Master) Get(ctx context.Context, srvName string) (*NodeInfo, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if node, ok := m.srvNodes[srvName]; ok {
		return &node, nil
	} else {
		return nil, errors.New("no this service nodes")
	}
}

//putEvent watch put event
func (m *Master) putEvent(ev *clientv3.Event) {
	var (
		nodeInfo NodeInfo
	)

	if err := json.Unmarshal(ev.Kv.Value, &nodeInfo); err != nil {
		log.Fatal(err)
	}

	m.add(nodeInfo)
}

//delEvent watch delete event
func (m *Master) delEvent(ev *clientv3.Event) {
	srvName := string(ev.Kv.Value)
	if err := m.del(srvName); err != nil {
		log.Fatal(err.Error())
	}
}

//add add node into master allNodes map
func (m *Master) add(nodeInfo NodeInfo) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.srvNodes[nodeInfo.SrvName] = nodeInfo
}

//del delete node from master allNodes map
func (m *Master) del(srvName string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.srvNodes[srvName]; ok {
		delete(m.srvNodes, srvName)
	}
	return nil
}
