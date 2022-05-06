package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	"xframe/config"
	"xframe/docs"
	grpcMall "xframe/internal/access/grpc/proto/mall"
	"xframe/internal/access/http/middleware"
	"xframe/internal/access/http/router"
	"xframe/internal/core"
	"xframe/pkg"
)

var App = &cli.App{
	Version: fmt.Sprintf("%s|%s|%s|%s",
		runtime.GOOS, runtime.GOARCH, config.BuildVersion, config.BuildAt),
	Action: Run,
}

func Run(c *cli.Context) error {
	time.Local, _ = time.LoadLocation("Asia/Shanghai")
	docs.SwaggerInfo.BasePath = "/"

	// 加载 telemetry
	shutdown := pkg.InitTracer()
	defer shutdown()

	var httpServer *http.Server
	var grpcServer *grpc.Server
	container := GetContainer()
	// http 服务
	err := container.Invoke(func(s *core.HttpServer, g *core.GrpcServer) {
		grpcServer = g.Engine
		go func() {
			if config.Conf.Common.Debug {
				gin.SetMode(gin.DebugMode)
			} else {
				gin.SetMode(gin.ReleaseMode)
			}
			_ = s.Engine.SetTrustedProxies(config.Conf.Server.TrustedProxies)
			s.Engine.Use(middleware.Exception)
			router.InitRout(s)
			httpServer = &http.Server{
				Addr:    config.Conf.Server.Addr,
				Handler: s.Engine,
			}
			_ = httpServer.ListenAndServe()
		}()
		go func() {
			lis, err := net.Listen("tcp", config.Conf.GrpcServer.Addr)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
				return
			}
			grpcMall.RegisterMallServer(g.Engine, g.ServerMall)
			log.Printf("server listening at %v", lis.Addr())
			if err := g.Engine.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		}()
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	cronjob := InitCron()
	cronjob.Start()

	// 监听信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	_ = httpServer.Shutdown(context.Background())
	grpcServer.GracefulStop()
	ctx := cronjob.Stop()

	log.Println("Shutting down cron...")
	select {
	case <-time.After(10 * time.Second):
		log.Fatal("Cron forced to shutdown...")
	case <-ctx.Done():
		log.Println("Cron exiting...")
	}
	return nil
}
