package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"im/internal/pkg/logger"
	"im/internal/pkg/signalctx"

	"go.uber.org/dig"
)

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
func (s *Controller) Run(signalctx *signalctx.Context) {
	// Run
	for _, server := range s.servers {
		s.mx.Lock()
		go func() {
			if err := server.Run(signalctx); err != nil {
				s.in.Logger.Error(signalctx, fmt.Errorf("run error: %v \n", err))
			}
		}()
		s.mx.Unlock()
	}

	// 監聽關機
	c := <-signalctx.Shutdown()
	slog.Info(fmt.Sprintf("Server Shutdown, osSignal: %v\n", c))

	// 所有使用signalctx的排程陸續關閉
	signalctx.Cancel()

	s.mx.Lock()
	defer s.mx.Unlock()

	var isException bool
	// 執行所有執行序shutdown
	for _, server := range s.servers {
		signalctx.Increment()
		go func(server IServer) {
			defer signalctx.Decrement()
			if err := server.Shutdown(signalctx); err != nil {
				isException = true
			}
			fmt.Println("debug shutdown")
		}(server)
	}

	// 確認所有關閉或是timeout
	select {
	case <-signalctx.AllDone():
		fmt.Println("debug AllDone")
		break
	case <-time.Tick(time.Second * 30):
		slog.ErrorContext(signalctx, "Shutdown Timeout!")
	}

	slog.Info("Server exit!")

	if isException {
		os.Exit(1)
	} else {
		os.Exit(0)
	}

}
