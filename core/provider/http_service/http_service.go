package http_service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"xframe/core/application"
	"xframe/pkg/config"
)

type HttpService struct {
	e *gin.Engine
	h *http.Server
}

type Option func(e *gin.Engine)

type OptionGroup struct {
	dig.In

	Opts []Option `group:"middle"`
}

func New(opts OptionGroup) application.Service {
	gin.SetMode(config.Conf.HttpServer.Mode)
	e := gin.New()
	_ = e.SetTrustedProxies(config.Conf.HttpServer.TrustedProxies)

	for _, opt := range opts.Opts {
		opt(e)
	}
	e.Use(gin.Recovery())

	server := &HttpService{e: e}
	server.h = &http.Server{
		Addr:    config.Conf.HttpServer.Addr,
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
