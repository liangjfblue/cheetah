package service

import (
	"context"

	"github.com/liangjfblue/cheetah/app/service/worker/common/job"

	"google.golang.org/grpc/status"

	"github.com/liangjfblue/cheetah/common/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	v1 "github.com/liangjfblue/cheetah/app/service/worker/proto/v1"
)

type WorkerService struct {
}

//调度新任务
func (w *WorkerService) StartJob(ctx context.Context, in *v1.StartJobRequest, out *v1.StartJobRespond) error {

	return nil
}

//停止任务
func (w *WorkerService) StopJob(ctx context.Context, in *v1.StopJobRequest, out *v1.StopJobRespond) error {
	return nil
}

//重启任务
func (w *WorkerService) RestartJob(ctx context.Context, in *v1.RestartJobRequest, out *v1.RestartRespond) error {
	return nil
}

//上报任务进度
func (w *WorkerService) JobProgress(ctx context.Context, in *v1.JobProgressRequest, stream v1.Worker_JobProgressStream) error {
	var (
		err       error
		jobStatus int32 = job.StatusReady
	)

	if ctx.Err() == context.Canceled {
		return errors.Wrap(status.New(codes.Canceled, "Client cancelled, abandoning").Err(), "service worker")
	}

	defer stream.Close()

	//根据in.ID 查询出当前运行的任务
	if jobStatus != job.StatusRunning {
		logger.Info("service worker job is not running")

		if err = stream.Send(&v1.JobProgressRespond{
			Code:     0,
			Progress: -1,
			Status:   jobStatus,
			Msg:      errors.New("job is error").Error(),
		}); err != nil {
			logger.Error("service worker err:%s", err.Error())
			return errors.Wrap(err, " service worker")
		}

		return nil
	}

	for {
		//TODO 查看任务进度, 状态, 发送

		if jobStatus != job.StatusRunning {
			logger.Info("service worker job is not running")

			if err = stream.Send(&v1.JobProgressRespond{
				Code:     1,
				Progress: 100,
				Status:   jobStatus,
			}); err != nil {
				logger.Error("service worker err:%s", err.Error())
				return errors.Wrap(err, " service worker")
			}

			return nil
		}

		if err = stream.Send(&v1.JobProgressRespond{
			Code:     1,
			Progress: 1,
			Status:   jobStatus,
		}); err != nil {
			logger.Error("service worker err:%s", err.Error())
			return errors.Wrap(err, " service worker")
		}
	}
}
