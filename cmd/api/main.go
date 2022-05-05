package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
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
	handlerGoods "xframe/access/http/handler/goods"
	"xframe/access/http/router"
	"xframe/config"
	repoGoods "xframe/infrastructure/repository/goods"
	"xframe/pkg"
	serviceGoods "xframe/service/goods"
)

const (
	ArgConfig         = "config"     // 启动参数的配置项 名称
	ArgConfigFilename = "config.ini" // 启动参数的配置项 值
)

func Stack() *cli.App {
	configFile := ""
	buildInfo := fmt.Sprintf("%s|%s|%s|%s",
		runtime.GOOS, runtime.GOARCH, config.BuildVersion, config.BuildAt)
	return &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        ArgConfig,
				Value:       ArgConfigFilename,
				Destination: &configFile,
			},
		},
		Name:    config.Conf.Server.Name,
		Version: buildInfo,
		Usage:   fmt.Sprintf("./%s -%s=%s", config.Conf.Server.Name, ArgConfig, ArgConfigFilename),
		Action: func(c *cli.Context) error {
			time.Local, _ = time.LoadLocation("Asia/Shanghai")

			var httpServer *http.Server
			var grpcServer *grpc.Server
			container := makeContainer()
			// http 服务
			go func() {
				// 初始化服务，注册路由
				s := gin.New()
				_ = s.SetTrustedProxies([]string{"0.0.0.0"})
				router.InitRout(s, container)
				httpServer = &http.Server{
					Addr:    config.Conf.Server.Addr,
					Handler: s,
				}
				_ = httpServer.ListenAndServe()
			}()
			// grpc 服务
			go func() {
				lis, err := net.Listen("tcp", ":2222")
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
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			select {
			case <-sigs:
				log.Println("notify sigs")
				_ = httpServer.Shutdown(context.Background())
				grpcServer.GetServiceInfo()
				log.Println("http shutdown")
			}
			return nil
		},
	}
}

func makeContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(pkg.NewMysql)
	_ = container.Provide(pkg.NewRedis)
	_ = container.Provide(repoGoods.New)
	_ = container.Provide(serviceGoods.New)
	_ = container.Provide(handlerGoods.New)
	_ = container.Provide(server.New)
	return container
}
