package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Types
const (
	// Test 测试环境
	Test = "test"
	// Dev 开发环境
	Dev = "dev"
	// Pro 正式环境
	Pro = "pro"
)

// Conf 全局配置文件
var Conf *Config

// Config 配置文件
type Config struct {
	HTTP HTTP `yaml:"http"`

	DB DB
}

type conf struct {
	Types string
	Dev   Config
	Test  Config
	Pro   Config
}

// NewConfig 获取配置配置
func NewConfig(path string) (*Config, error) {
	if path == "" {
		path = "configs/config.yml"
	}

	var conf conf
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return nil, err
	}
	var config *Config
	switch conf.Types {
	case "test":
		config = &conf.Test
	case "dev":
		config = &conf.Dev
	case "pro":
		config = &conf.Pro
	default:
		return nil, errors.New("unknown environment configuration")
	}

	Conf = config
	return config, nil
}

// HTTP http 配置
type HTTP struct {
	Addr string
}

// DB 存储配置
type DB struct {
	Postgres Postgres
	Redis    Redis
}

// Redis 存储配置
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Postgres 配置
type Postgres struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB   string `yaml:"db"`

	User     string `yaml:"user"`
	Password string `yaml:"password"`

	SSLMode string `yaml:"sslMode"`

	MaxIdle int `yaml:"maxIdle"`
	MaxOpen int `yaml:"maxOpen"`
}

func (p *Postgres) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		p.Host, p.Port, p.User, p.DB, p.Password, p.SSLMode)
}
