package config

import "time"

var Cfg = &Config{}

type Config struct {
	Text   string     `yaml:"test" json:"test"`     //
	DB     *DbItem    `yaml:"db" json:"db"`         //
	Redis  *RedisItem `yaml:"redis" json:"redis"`   //
	Zipkin *Zipkin    `yaml:"zipkin" json:"zipkin"` //
}

type Zipkin struct {
	Url string `yaml:"url" json:"url"` //
}

type DbItem struct {
	Name            string        `yaml:"name" json:"name"`                             //
	Type            string        `yaml:"type" json:"type"`                             //
	Server          string        `yaml:"server" json:"server"`                         //
	Port            int           `yaml:"port" json:"port"`                             //
	Database        string        `yaml:"database" json:"database"`                     //
	User            string        `yaml:"user" json:"user"`                             //
	Password        string        `yaml:"password" json:"password"`                     //
	ConnMaxLifeTime time.Duration `yaml:"conn_max_life_time" json:"conn_max_life_time"` //
	MaxIdleConn     int           `yaml:"max_idle_conn" json:"max_idle_conn"`           //
	MaxOpenConn     int           `yaml:"max_open_conn" json:"max_open_conn"`           //
}

type RedisItem struct {
	Name string `yaml:"name" json:"name"` //
	Addr string `yaml:"addr" json:"addr"` //
	Auth string `yaml:"auth" json:"auth"` //
	DB   int    `yaml:"db" json:"db"`     //
}
