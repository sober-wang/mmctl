package types

// AppID 信息
type AppID struct {
	Name      string `josn:"name"`
	CPUs      string `josn:"cpus"`
	Mem       string `josn:"mem"`
	Running   string `josn:"running"`
	Staged    string `josn:"staged"`
	Health    string `josn:"health"`
	Unhealthy string `josn:"unHealthy"`
}

// ConfPath 配置文件路径
var ConfPath = "conf/mmctl.yaml"

// FrameWork mesos framework 信息
type FrameWork struct {
	Name     string `json:"name"`
	HostName string `json:"hostname"`
	WebeUi   string `json:"webeui_url"`
	Active   string `json:"active"`
}

// NodeInfo node 信息
type NodeInfo struct {
	Name       string `json:"name"`
	Attributes string `json:"attributes"`
	IP         string `json:"ip"`
	Active     string `json:"active"`
	Version    string `json:"version"`
	AllCPU     string `json:"allCPUs"`
	AllMem     string `json:"allMem"`
	AllDisk    string `json:"allDisk"`
	UsedCPU    string `json:"usedCPU"`
	UsedMem    string `json:"usedMem"`
	UsedDisk   string `json:"UsedDisk"`
}
