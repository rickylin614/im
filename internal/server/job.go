package server

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"go.uber.org/dig"

	"im/internal/listener"
	"im/internal/pkg/config"
	"im/internal/pkg/logger"
)

type JobDigIn struct {
	dig.In

	Config *config.Config
	Logger logger.Logger

	// Job

	// IListener
	MessageListener     listener.IListener `name:"messageListener"`
	MessageSaveListener listener.IListener `name:"messageSaveListener"`
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
	s.test()
	s.StartListener()

	s.job.StartBlocking()
	return nil
}

func (s *JobServer) Shutdown(ctx context.Context) error {
	s.job.Stop()
	return nil
}

func (s *JobServer) StartJob(in JobDigIn) {
	// Job config
}

func (s *JobServer) StartListener() {
	c := s.In.Config.ListenerConfig
	if !c.Enable {
		return
	}
	if c.Msg > 0 {
		s.In.MessageListener.Start(c.Msg)
	}
	if c.MsgSave > 0 {
		s.In.MessageSaveListener.Start(c.MsgSave)
	}
}

func (s *JobServer) test() {
	i := 0
	s.job.Every("3600s").Do(func(i *int) {
		*i++
		data := *i
		fmt.Println("Job exec start", data)
		time.Sleep(time.Second * 10)
		fmt.Println("Job exec end", data)
	}, &i)
}

//func (s *JobServer) Run(ctx context.Context) error {
//	defer func() {
//		if err := recover(); err != nil {
//			s.Run(ctx)
//		}
//	}()
//
//	logger := CronLogger{Logger: s.In.Logger}
//	taipeiLocation, _ := time.LoadLocation("Asia/Taipei")
//	s.job = cron.New(
//		cron.WithChain(
//			cron.Recover(logger),
//			cron.SkipIfStillRunning(logger),
//		),                                 // 有需要其他中間件可再添加
//		cron.WithLocation(taipeiLocation), // UTC+8 時區
//		cron.WithLogger(logger),
//	)
//
//	i := 0
//	s.job.AddFunc("*/2 * * * * *", func() {
//		i++
//		data := i
//		fmt.Println("Job exec start", data)
//		time.Sleep(time.Second * 10)
//		fmt.Println("Job exec end", data)
//		panic(errors.New("panic test"))
//	})
//
//	return nil
//}
//
//
//type CronLogger struct {
//	Logger logger.Logger
//}
//
//func (c CronLogger) Info(msg string, keysAndValues ...interface{}) {
//	m := make(map[any]any, len(keysAndValues)/2)
//	for i := 0; i+1 < len(keysAndValues); i += 2 {
//		m[keysAndValues[i]] = keysAndValues[i+1]
//	}
//	message := fmt.Sprintf("message: %s , datas: %+v", msg, m)
//	c.Logger.Info(context.Background(), message)
//}
//
//func (c CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
//	m := make(map[any]any, len(keysAndValues)/2)
//	for i := 0; i+1 < len(keysAndValues); i += 2 {
//		m[keysAndValues[i]] = keysAndValues[i+1]
//	}
//	msgErr := fmt.Errorf("err: %w, message: %s , datas: %+v", err, msg, m)
//	c.Logger.Error(context.Background(), msgErr)
//}
