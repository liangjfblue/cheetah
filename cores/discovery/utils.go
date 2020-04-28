package discovery

import (
	"encoding/json"
	"path"
	"strings"
)

var (
	prefix = "/discovery/services/"
)

func Encode(s *Service) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func Decode(ds []byte) *Service {
	var s *Service
	json.Unmarshal(ds, &s)
	return s
}

//nodePath 以防node中有/误导
func NodePath(s, id string) string {
	service := strings.Replace(s, "/", "-", -1)
	node := strings.Replace(id, "/", "-", -1)
	return path.Join(prefix, service, node)
}

//servicePath 以防srvName中有/误导
func ServicePath(s string) string {
	return path.Join(prefix, strings.Replace(s, "/", "-", -1))
}

//ServicePrefixPath 所有服务根目录
func ServicePrefixPath() string {
	return prefix
}

//Copy 拷贝服务
func Copy(current []*Service) []*Service {
	services := make([]*Service, len(current))
	for i, service := range current {
		services[i] = CopyService(service)
	}
	return services
}

//CopyService 深拷贝服务
func CopyService(service *Service) *Service {
	s := new(Service)
	*s = *service

	nodes := make([]*Node, len(service.Nodes))
	for j, node := range service.Nodes {
		n := new(Node)
		*n = *node
		nodes[j] = n
	}
	s.Nodes = nodes

	eps := make([]*Endpoint, len(service.Endpoints))
	for j, ep := range service.Endpoints {
		e := new(Endpoint)
		*e = *ep
		eps[j] = e
	}
	s.Endpoints = eps
	return s
}
