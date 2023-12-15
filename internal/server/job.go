package server

import (
	"context"
	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"time"

	"github.com/go-co-op/gocron"
	"go.uber.org/dig"
)

type JobDigIn struct {
	dig.In

	Config       *config.Config
	Logger       logger.Logger
	ServerRunner *Controller
}

type JobServer struct {
	job *gocron.Scheduler

	In JobDigIn
}

func (s *JobServer) Run(ctx context.Context) error {
	s.job = gocron.NewScheduler(time.UTC) // UTC +0

	s.job.StartBlocking()
	return nil
}

func (s *JobServer) Shutdown(ctx context.Context) error {
	done := make(chan struct{})

	go func() {
		s.job.Stop()
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}
