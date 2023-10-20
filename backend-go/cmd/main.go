package main

import (
	"backend-go/internal/pkg/config"
)

var conf config.Conf

func main() {
	// app.InitRouter()

	conf.GetConf()
}
