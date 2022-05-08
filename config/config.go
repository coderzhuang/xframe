package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

var (
	BuildVersion string // 编译的app版本
	BuildAt      string // 编译时间
)

var Conf = Config{}

type Common struct {
	Debug bool `yaml:"debug" json:"debug"` // debug
}

type HttpServer struct {
	Switch         bool     `yaml:"switch" json:"switch"`                   // 开关
	Name           string   `yaml:"name" json:"name"`                       // 服务名称
	Addr           string   `yaml:"addr" json:"addr"`                       // 服务地址
	Mode           string   `yaml:"mode" json:"mode"`                       // gin Mode
	TrustedProxies []string `yaml:"trusted_proxies" json:"trusted_proxies"` // 信任的代理
}

type GrpcServer struct {
	Switch bool   `yaml:"switch" json:"switch"` // 开关
	Name   string `yaml:"name" json:"name"`     // 服务名称
	Addr   string `yaml:"addr" json:"addr"`     // 服务地址
}

type CronServer struct {
	Switch bool `yaml:"switch" json:"switch"` // 开关
}

type Swagger struct {
	Switch bool `yaml:"switch" json:"switch"` // 开关
}

type Zipkin struct {
	Url string `yaml:"url" json:"url"` //
}
type DbItem struct {
	Name     string `yaml:"name" json:"name"`         //
	Type     string `yaml:"type" json:"type"`         //
	Server   string `yaml:"server" json:"server"`     //
	Port     int    `yaml:"port" json:"port"`         //
	Database string `yaml:"database" json:"database"` //
	User     string `yaml:"user" json:"user"`         //
	Password string `yaml:"password" json:"password"` //
}

type RedisItem struct {
	Name string `yaml:"name" json:"name"` //
	Addr string `yaml:"addr" json:"addr"` //
	Auth string `yaml:"auth" json:"auth"` //
	DB   int    `yaml:"db" json:"db"`     //
}

type Config struct {
	Common     *Common     `yaml:"common" json:"common"`           //
	HttpServer *HttpServer `yaml:"http_server" json:"http_server"` //
	GrpcServer *GrpcServer `yaml:"grpc_server" json:"grpc_server"` //
	CronServer *CronServer `yaml:"cron_server" json:"cron_server"` //
	Swagger    *Swagger    `yaml:"swagger" json:"swagger"`         //
	DB         *DbItem     `yaml:"db" json:"db"`                   //
	Redis      *RedisItem  `yaml:"redis" json:"redis"`             //
	Zipkin     *Zipkin     `yaml:"zipkin" json:"zipkin"`           //
}

func init() {
	configFile := "./config/config.yaml"
	_, err := os.Stat(configFile)
	if !(err == nil || os.IsExist(err)) {
		panic("config file does not exists")
	}
	b, _ := ioutil.ReadFile(configFile)
	_ = yaml.Unmarshal(b, &Conf)
}
