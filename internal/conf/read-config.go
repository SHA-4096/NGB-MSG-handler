package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	LogConfig logConfigStruct `yaml:"LogConfig"`
}

type logConfigStruct struct {
	LogPath    string `yaml:"logPath"`
	RotateTime int    `yaml:"rotateTime"`
	MaxAge     int    `yaml:"maxAge"`
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
