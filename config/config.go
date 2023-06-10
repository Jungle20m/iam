package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type AppConfig struct {
	AppName     string `yaml:"app_name"`
	ServiceName string `yaml:"service_name"`
	Language    string `yaml:"language"`
	Environment string `yaml:"environment"`
	Debug       bool   `yaml:"debug"`
}

type ApiConfig struct {
	HttpHost string `yaml:"http_host"`
	HttpPort string `yaml:"http_port"`
	GrpcHost string `yaml:"grpc_host"`
	GrpcPort string `yaml:"grpc_port"`
}

type SwaggerConfig struct {
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
	Host    string `yaml:"host"`
}

type MysqlConfig struct {
	Dsn string `yaml:"dsn"`
}

type LogConfig struct {
	Output    string `yaml:"output"`
	Formatter string `yaml:"formatter"`
	Level     string `yaml:"level"`
}

type Config struct {
	App     AppConfig     `yaml:"app"`
	Api     ApiConfig     `yaml:"api"`
	Swagger SwaggerConfig `yaml:"swagger"`
	Mysql   MysqlConfig   `yaml:"mysql"`
	Log     LogConfig     `yaml:"log"`
}

type Tracer struct {
}

func LoadConfig() (*Config, error) {
	confData, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("loading config file failed: %v", err)
	}
	conf := Config{}
	err = yaml.Unmarshal(confData, &conf)
	if err != nil {
		return nil, fmt.Errorf("reading config failed: %v", err)
	}
	return &conf, nil
}
