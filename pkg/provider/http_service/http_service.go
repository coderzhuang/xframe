package http_service

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"xframe/config"
	"xframe/pkg/application"
)

type HttpService struct {
	e *gin.Engine
	h *http.Server
}

func New(e *gin.Engine) application.Service {
	server := &HttpService{e: e}
	server.h = &http.Server{
		Addr:    config.Conf.Server.Addr,
		Handler: e,
	}
	return server
}

func (s *HttpService) Run() {
	go func() {
		_ = s.h.ListenAndServe()
	}()
}

func (s *HttpService) Shutdown() {
	_ = s.h.Shutdown(context.Background())
}
