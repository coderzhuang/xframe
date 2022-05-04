package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"log"
	"net"
	"runtime"
	grpcMall "test/access/grpc/proto/mall"
	"test/access/grpc/server"
	handlerGoods "test/access/http/handler/goods"
	"test/access/http/router"
	"test/config"
	repoGoods "test/infrastructure/repository/goods"
	"test/pkg"
	serviceGoods "test/service/goods"
	"time"
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

			var err error
			container := dig.New()
			if err = container.Provide(pkg.NewMysql); err != nil {
				fmt.Println(err.Error())
			}
			if err = container.Provide(pkg.NewRedis); err != nil {
				fmt.Println(err.Error())
			}
			if err = container.Provide(repoGoods.New); err != nil {
				fmt.Println(err.Error())
			}
			if err = container.Provide(serviceGoods.New); err != nil {
				fmt.Println(err.Error())
			}
			if err = container.Provide(handlerGoods.New); err != nil {
				fmt.Println(err.Error())
			}
			if err = container.Provide(server.New); err != nil {
				fmt.Println(err.Error())
			}

			// http 服务
			go func() {
				// 初始化服务，注册路由
				s := gin.New()
				_ = s.SetTrustedProxies([]string{"0.0.0.0"})
				router.InitRout(s, container)
				_ = s.Run(config.Conf.Server.Addr)
			}()
			// grpc 服务
			go func() {
				lis, err := net.Listen("tcp", ":2222")
				if err != nil {
					log.Fatalf("failed to listen: %v", err)
				}
				s := grpc.NewServer()

				err = container.Invoke(func(service *server.Mall) {
					grpcMall.RegisterMallServer(s, service)
					log.Printf("server listening at %v", lis.Addr())
					if err := s.Serve(lis); err != nil {
						log.Fatalf("failed to serve: %v", err)
					}
				})
				if err != nil {
					fmt.Println(err.Error())
				}
			}()
			return nil
		},
	}
}
