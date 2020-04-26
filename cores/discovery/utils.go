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
