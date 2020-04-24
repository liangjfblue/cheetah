package discovery

// NodeInfo node info
type NodeInfo struct {
	Version  string `json:"version"`  //版本号
	Path     string `json:"path"`     //注册路径
	Env      string `json:"env"`      //环境
	SrvName  string `json:"srvName"`  //服务名
	Addr     string `json:"addr"`     //服务地址
	Hostname string `json:"hostname"` //主机名（必须唯一）
	Status   int    `json:"status"`   //状态，1表示接收流量，2表示不接收
	Color    string `json:"color"`    //灰度或集群标识
}

// Member master node info
type Member struct {
	Online bool //在线标志
	NodeInfo
}
