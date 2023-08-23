package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	LogConfig  logConfigStruct  `yaml:"LogConfig"`
	AMQPConfig AMQPConfigStruct `yaml:"AMQPConfig"`
	Database   DatabaseStruct   `yaml:"Database"`
}

type logConfigStruct struct {
	LogPath    string `yaml:"logPath"`
	RotateTime int    `yaml:"rotateTime"`
	MaxAge     int    `yaml:"maxAge"`
}

type AMQPConfigStruct struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type DatabaseStruct struct {
	Mysql MysqlStruct `yaml:"Mysql"`
}

type MysqlStruct struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbName"`
	Timeout  string `yaml:"timeout"`
}

var Config *ConfigStruct

func ReadConfig() {
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic("Error when loading config")
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic("Error when unmarshaling")
	}
	fmt.Println("Config loaded")
}
