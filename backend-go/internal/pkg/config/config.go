package config

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type database struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type Conf struct {
	DataBase database `yaml:"database"`
}

func (c *Conf) GetConf() {
	os_wd_path, _ := os.Getwd()
	path := path.Join(os_wd_path, "config/app.yaml")

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("read config file failed, err: ", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println("parse config file err: ", err)
		os.Exit(1)
	}
}
