package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type AppConfig struct {
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
	Swagger SwaggerConfig `yaml:"swagger"`
	Mysql   MysqlConfig   `yaml:"mysql"`
	Log     LogConfig     `yaml:"log"`
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
