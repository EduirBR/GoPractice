package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Port          string `yaml:"PORT"`
	HostDB        string `yaml:"HOSTPOSTGRES"`
	PortPosgresql string `yaml:"PORTPOSTGRES"`
	DBName        string `yaml:"DBPOSTGRES"`
	User          string `yaml:"USUPOSTGRES"`
	Password      string `yaml:"PASSPOSTGRES"`
}

//GetConfig carga las configuraciones del archivo Config.yaml
//retorna las configuraciones contenidas en el archivo y el err si existe
func GetConfig() (conf *Settings, err error) {
	data, err := ioutil.ReadFile("../config/config.yaml")
	if err != nil {
		return
	}
	conf = &Settings{}
	err = yaml.Unmarshal(data, conf)
	return
}
