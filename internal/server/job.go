package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/pkg/signalctx"

	"github.com/go-co-op/gocron"
	"go.uber.org/dig"
)

type JobDigIn struct {
	dig.In

	Config       *config.Config
	Logger       logger.Logger
	ServerRunner *Controller
	Ctx          *signalctx.Context
}

type JobServer struct {
	job *gocron.Scheduler

	In JobDigIn
}

func (s *JobServer) Run(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			s.Run(ctx)
		}
	}()

	s.job = gocron.NewScheduler(time.UTC) // UTC +0
	gocron.SetPanicHandler(func(jobName string, recoverData interface{}) {
		msg := "job panic! name:" + jobName
		s.In.Logger.Error(context.Background(), fmt.Errorf("%s %v", msg, recoverData))
	})

	i := 0
	s.job.Every("1s").Do(func(i *int) {
		*i++
		data := *i
		fmt.Println("Job exec start", data)
		time.Sleep(time.Second * 10)
		fmt.Println("Job exec end", data)
		panic(errors.New("panic test"))
	}, &i)

	s.job.StartBlocking()
	return nil
}

func (s *JobServer) Shutdown(ctx context.Context) error {
	s.job.Stop()
	return nil
}
