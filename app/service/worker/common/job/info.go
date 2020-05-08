package job

type JobIDType int32

//JobInfo 任务结构体
type JobInfo struct {
	ID   JobIDType `json:"id"`
	Name string    `json:"name"`
	Type int32     `json:"type"`
	Code string    `json:"code"`
}
