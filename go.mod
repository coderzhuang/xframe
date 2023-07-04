module xframe

go 1.16

require (
	github.com/coderzhuang/core v1.0.2
	github.com/gin-gonic/gin v1.9.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/jinzhu/copier v0.3.5
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.9.3
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.1
	go.uber.org/dig v1.15.0
	google.golang.org/grpc v1.56.1
	google.golang.org/protobuf v1.30.0
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.2
)

replace github.com/coderzhuang/core => ../x-core
