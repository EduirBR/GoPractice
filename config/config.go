package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Port          string `yaml:"PORT"`
	HostDB        string `yaml:"HOSTPOSTGRES"`
	PortPosgresql string `yaml:"PORTPOSTGRES"`
	DBName        string `yaml:"DBPOSTGRES"`
	User          string `yaml:"USUPOSTGRES"`
	Password      string `yaml:"PASSPOSTGRES"`
	FileBaseName  string `yaml:"FILEBASENAME"`
	Creator       string `yaml:"CREATOR"`
}

func GetConfig() (conf *Settings, err error) {
	data, err := os.ReadFile("../config/config.yaml")
	if err != nil {
		return
	}
	conf = &Settings{}
	err = yaml.Unmarshal(data, conf)
	return
}
