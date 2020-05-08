package memory

import (
	"sync"

	"github.com/pkg/errors"

	"github.com/liangjfblue/cheetah/app/service/worker/service/cache"

	"github.com/liangjfblue/cheetah/app/service/worker/common/job"
)

var (
	ErrJobIsNil    = errors.New("job is nil")
	ErrJobNotExist = errors.New("job is not exist")
	ErrJobIsEmpty  = errors.New("job is empty")
)

type jobCache struct {
	Opts          cache.Options
	allJobMap     map[job.JobIDType]*job.JobInfo
	runningJobMap map[job.JobIDType]*job.JobInfo
	lock          sync.RWMutex
}

//初始化cache
func (c *jobCache) Init(opts ...cache.Option) {
	for _, o := range opts {
		o(&c.Opts)
	}
}

//获取参数结构体
func (c *jobCache) Options() cache.Options {
	return c.Opts
}

//新增任务到所有缓存map
func (c *jobCache) Add2All(job *job.JobInfo) error {
	if job == nil {
		return ErrJobIsNil
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	c.allJobMap[job.ID] = job

	return nil
}

//新增任务到所有缓存map
func (c *jobCache) Add2Running(job *job.JobInfo) error {
	if job == nil {
		return ErrJobIsNil
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	c.runningJobMap[job.ID] = job

	return nil
}

//删除缓存任务
func (c *jobCache) Delete(jobID job.JobIDType) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.allJobMap[jobID]; ok {
		delete(c.allJobMap, jobID)
	}

	if _, ok := c.runningJobMap[jobID]; ok {
		delete(c.runningJobMap, jobID)
	}

	return nil
}

//获取任务
func (c *jobCache) Get(jobID job.JobIDType) (*job.JobInfo, error) {
	c.lock.RLock()
	defer c.lock.Unlock()

	if j, ok := c.runningJobMap[jobID]; ok {
		return j, nil
	}

	return nil, ErrJobNotExist
}

//获取所有任务
func (c *jobCache) All() ([]job.JobInfo, error) {
	c.lock.RLock()
	defer c.lock.Unlock()

	if len(c.allJobMap) <= 0 {
		return nil, ErrJobIsEmpty
	}

	jobs := make([]job.JobInfo, 0, len(c.allJobMap))
	for _, jobInfo := range c.allJobMap {
		jobs = append(jobs, *jobInfo)
	}
	return jobs, nil
}

//获取所有运行任务
func (c *jobCache) AllRunning() ([]job.JobInfo, error) {
	c.lock.RLock()
	defer c.lock.Unlock()

	if len(c.allJobMap) <= 0 {
		return nil, ErrJobIsEmpty
	}

	jobs := make([]job.JobInfo, 0, len(c.runningJobMap))
	for _, jobInfo := range c.runningJobMap {
		jobs = append(jobs, *jobInfo)
	}

	return jobs, nil
}

func New(opts ...cache.Option) cache.ICache {
	c := new(jobCache)
	c.Opts = cache.DefaultOptions
	c.allJobMap = make(map[job.JobIDType]*job.JobInfo)
	c.runningJobMap = make(map[job.JobIDType]*job.JobInfo)

	for _, o := range opts {
		o(&c.Opts)
	}

	return c
}
