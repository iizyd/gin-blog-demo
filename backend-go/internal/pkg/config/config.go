package config

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type server struct {
	RunMode      string `yaml:"run_mode"`
	HttpPort     string `yaml:"http_port"`
	ReadTimeout  string `yaml:"read_timeout"`
	WriteTimeout string `yaml:"write_timeout"`
}

type app struct {
	DefaultPageSize      int      `yaml:"default_page_size"`
	MaxPageSize          int      `yaml:"max_page_size"`
	LogSavePath          string   `yaml:"log_save_path"`
	LogFileName          string   `yaml:"log_file_name"`
	LogFileExt           string   `yaml:"log_file_ext"`
	UploadSavePath       string   `yaml:"upload_save_path"`
	UploadServerUrl      string   `yaml:"upload_server_url"`
	UploadImageMaxSize   int      `yaml:"upload_image_max_size"`
	UploadImageAllowExts []string `yaml:"upload_image_allow_exts"`
}

type database struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DB_Name  string `yaml:"db_name"`
}

type jwt struct {
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
	Expire int    `yaml:"expire"`
}

type Conf struct {
	Server   server   `yaml:"server"`
	App      app      `yaml:"app"`
	DataBase database `yaml:"database"`
	Jwt      jwt      `yaml:"jwt"`
}

func (c *Conf) getConf() {
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

var Config Conf

func InitConf() {
	Config.getConf()
}
