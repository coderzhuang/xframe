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
	grpcMall "xframe/access/grpc/proto/mall"
	"xframe/access/grpc/server"
	"xframe/access/http/middleware"
	"xframe/access/http/router"
	"xframe/config"
	"xframe/docs"
	"xframe/pkg"
)

var App = &cli.App{
	Version: fmt.Sprintf("%s|%s|%s|%s",
		runtime.GOOS, runtime.GOARCH, config.BuildVersion, config.BuildAt),
	Action: Run,
}

func Run(c *cli.Context) error {
	time.Local, _ = time.LoadLocation("Asia/Shanghai")
	docs.SwaggerInfo_swagger.BasePath = "/"

	// 加载 telemetry
	shutdown := pkg.InitTracer()
	defer shutdown()

	var httpServer *http.Server
	var grpcServer *grpc.Server
	container := GetContainer()
	// http 服务
	go func() {
		// 初始化服务，注册路由
		s := gin.New()
		_ = s.SetTrustedProxies([]string{"0.0.0.0"})
		s.Use(middleware.Exception)
		router.InitRout(s, container)
		httpServer = &http.Server{
			Addr:    config.Conf.Server.Addr,
			Handler: s,
		}
		_ = httpServer.ListenAndServe()
	}()
	// grpc 服务
	go func() {
		lis, err := net.Listen("tcp", config.Conf.GrpcServer.Addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer = grpc.NewServer()
		err = container.Invoke(func(service *server.Mall) {
			grpcMall.RegisterMallServer(grpcServer, service)
			log.Printf("server listening at %v", lis.Addr())
			if err := grpcServer.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	// 定时脚本
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
