package cache

import "github.com/liangjfblue/cheetah/app/service/worker/common/job"

type ICache interface {
	//初始化cache
	Init(...Option)
	//获取参数结构体
	Options() Options
	//新增任务到所有缓存map
	Add2All(*job.JobInfo) error
	//新增任务到运行缓存map
	Add2Running(*job.JobInfo) error
	//删除缓存任务
	Delete(job.JobIDType) error
	//获取任务
	Get(job.JobIDType) (*job.JobInfo, error)
	//获取所有任务
	All() ([]job.JobInfo, error)
	//获取所有运行任务
	AllRunning() ([]job.JobInfo, error)
}

type Option func(opts *Options)
