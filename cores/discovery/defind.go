package discovery

// NodeInfo node info
type Service struct {
	SrvName   string            `json:"srvName"`   //服务名
	Version   string            `json:"version"`   //版本号
	Metadata  map[string]string `json:"metadata"`  //服务元数据
	Endpoints []*Endpoint       `json:"endpoints"` //
	Nodes     []*Node           `json:"nodes"`
}

type Node struct {
	Id       string            `json:"id"`
	Address  string            `json:"address"`
	Metadata map[string]string `json:"metadata"`
}

type Endpoint struct {
	Name     string            `json:"name"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	Metadata map[string]string `json:"metadata"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values"`
}
