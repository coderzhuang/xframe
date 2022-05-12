package mysql

import (
	"fmt"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"xframe/pkg/config"
)

var db *gorm.DB

func New() *gorm.DB {
	if db != nil {
		return db
	}

	logLevel := logger.Error
	if config.Conf.Common.Debug {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		//GetLog(),// mysql的错误等级会转换成log的info，底层调用的是log.Printf，当log设置error后，mysql的error信息也不能输出了
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	conf := config.Conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Server, conf.Port, conf.Database,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(config.Conf.DB.MaxOpenConn)
	sqlDB.SetMaxIdleConns(config.Conf.DB.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Minute * config.Conf.DB.ConnMaxLifeTime)

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
	return db
}
