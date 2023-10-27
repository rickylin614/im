package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/router"

	"go.uber.org/dig"
)

// digIn the dependency invoke Server
type digIn struct {
	dig.In

	WebRouter    *router.WebRouter
	Config       *config.Config
	Logger       logger.Logger
	ServerRunner *Controller
}

type serverDigIn struct {
	dig.In

	Logger logger.Logger
}

type IServer interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type Controller struct {
	mx      sync.Mutex
	servers []IServer

	in serverDigIn
}

func NewServerController(in serverDigIn) *Controller {
	return &Controller{in: in}
}

func (s *Controller) Register(server IServer) {
	s.mx.Lock()
	s.servers = append(s.servers, server)
	s.mx.Unlock()
}

// Run Listen 監聽服務, 若確認所有服務正常關閉則os.Exit
func (s *Controller) Run() {
	// Run
	ctx := context.Background()
	for _, server := range s.servers {
		s.mx.Lock()
		go func() {
			if err := server.Run(ctx); err != nil {
				s.in.Logger.Error(ctx, fmt.Errorf("run error: %v \n", err))
			}
		}()
		s.mx.Unlock()
	}

	// listen and graceful shutdown
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, os.Kill)
	c := <-shutdownSignal

	s.mx.Lock()
	defer s.mx.Unlock()
	fmt.Printf("Server Shutdown, osSignal: %v\n", c)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	errs := make(chan error, len(s.servers))
	var wg sync.WaitGroup
	wg.Add(len(s.servers))
	for _, server := range s.servers {
		go func(server IServer) {
			defer wg.Done()
			errs <- server.Shutdown(ctx)
		}(server)
	}

	wg.Wait()
	close(errs)

	var isException bool

	for err := range errs {
		if err != nil {
			isException = true
			s.in.Logger.Error(context.TODO(), fmt.Errorf("shutdown error: %w", err))
		}
	}

	s.in.Logger.Info(context.TODO(), "Server exiting")

	if isException {
		os.Exit(1)
	} else {
		os.Exit(0)
	}

}
