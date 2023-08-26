package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
	srv *http.Server

	In digIn
}

func (s *WebServer) Run(context.Context) error {
	r := gin.New()

	s.In.WebRouter.SetRouter(r)

	s.srv = &http.Server{
		Addr:    s.In.Config.GinConfig.Port,
		Handler: r,
	}

	if err := s.srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("listen: %s\n", err)
		return err
	}

	return nil
}

func (s *WebServer) Shutdown(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("Server Shutdown: %w\n", err)
	}

	return nil
}
