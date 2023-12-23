package server

import (
	"context"
	"fmt"
	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/pkg/signalctx"
	"time"

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
	s.job = gocron.NewScheduler(time.UTC) // UTC +0
	i := 0
	s.job.Every("1s").Do(func(i *int) {
		*i++
		data := *i
		fmt.Println("Job exec start", data)
		time.Sleep(time.Second * 10)
		fmt.Println("Job exec end", data)
	}, &i)

	s.job.StartBlocking()
	return nil
}

func (s *JobServer) Shutdown(ctx context.Context) error {
	s.job.Stop()
	return nil
}
