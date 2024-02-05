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
	Ctx    *signalctx.Context
}

type IServer interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type SrvCtrl struct {
	mx      sync.Mutex
	servers []IServer

	in serverDigIn
}

func NewServerController(in serverDigIn) *SrvCtrl {
	return &SrvCtrl{in: in}
}

func (s *SrvCtrl) Register(server IServer) {
	s.mx.Lock()
	s.servers = append(s.servers, server)
	s.mx.Unlock()
}

// Run Listen 監聽服務, 若確認所有服務正常關閉則os.Exit
func (s *SrvCtrl) Run() {
	ctx := s.in.Ctx
	// Run
	for _, server := range s.servers {
		s.mx.Lock()
		go func(server IServer) {
			if err := server.Run(ctx); err != nil {
				s.in.Logger.Error(ctx, fmt.Errorf("run error: %v \n", err))
			}
		}(server)
		s.mx.Unlock()
	}

	// 監聽關機
	c := <-ctx.Shutdown()
	slog.Info(fmt.Sprintf("Server Shutdown, osSignal: %v\n", c))

	// 所有使用signalctx的排程陸續關閉
	ctx.Cancel()

	s.mx.Lock()
	defer s.mx.Unlock()

	var isException bool
	// 執行所有執行序shutdown
	for _, server := range s.servers {
		ctx.Increment()
		go func(server IServer) {
			defer ctx.Decrement()
			if err := server.Shutdown(ctx); err != nil {
				isException = true
				slog.Error("server.Shutdown err", "error", err.Error())
			}
			slog.Debug("debug shutdown")
		}(server)
	}

	// 確認所有關閉或是timeout
	select {
	case <-ctx.AllDone():
		slog.Debug("debug AllDone")
		break
	case <-time.Tick(time.Second * 30):
		slog.ErrorContext(ctx, "Shutdown Timeout!")
	}

	slog.Info("Server exit!")

	if isException {
		os.Exit(1)
	} else {
		os.Exit(0)
	}

}
