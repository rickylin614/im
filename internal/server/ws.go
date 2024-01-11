package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"im/internal/pkg/config"
	"im/internal/pkg/logger"
	"im/internal/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type WsDigIn struct {
	dig.In

	WsRouter *router.WsRouter
	Config   *config.Config
	Logger   logger.Logger
}

type WsServer struct {
	srv *http.Server

	In WsDigIn
}

func (s *WsServer) Run(context.Context) error {
	r := gin.New()

	s.In.WsRouter.SetRouter(r)

	s.srv = &http.Server{
		Addr:    s.In.Config.WsConfig.Port,
		Handler: r,
	}

	if err := s.srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("listen: %s\n", err)
		return err
	}

	return nil
}

func (s *WsServer) Shutdown(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("Server Shutdown: %w\n", err)
	}

	return nil
}
