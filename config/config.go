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

type Server struct {
	Name  string `yaml:"name" json:"name"`   // 服务名称
	Addr  string `yaml:"addr" json:"addr"`   // 服务地址
	Debug bool   `yaml:"debug" json:"debug"` // debug
}

type Zipkin struct {
	Url string `yaml:"url" json:"url"`
}
type DbItem struct {
	Name     string `yaml:"name" json:"name"`
	Type     string `yaml:"type" json:"type"`
	Server   string `yaml:"server" json:"server"`
	Port     int    `yaml:"port" json:"port"`
	Database string `yaml:"database" json:"database"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
}

type RedisItem struct {
	Name string `yaml:"name" json:"name"`
	Addr string `yaml:"addr" json:"addr"`
	Auth string `yaml:"auth" json:"auth"`
	DB   int    `yaml:"db" json:"db"`
}

type Config struct {
	Server *Server    `yaml:"server" json:"server"`
	DB     *DbItem    `yaml:"db" json:"db"`
	Redis  *RedisItem `yaml:"redis" json:"redis"`
	Zipkin *Zipkin    `yaml:"zipkin" json:"zipkin"`
}

func Init() {
	configFile := "./config/config.yaml"
	_, err := os.Stat(configFile)
	if !(err == nil || os.IsExist(err)) {
		panic("config file does not exists")
	}
	b, _ := ioutil.ReadFile(configFile)
	_ = yaml.Unmarshal(b, &Conf)
}
