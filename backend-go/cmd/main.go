package main

import (
	"backend-go/internal/app"
	"backend-go/internal/pkg/config"
	"backend-go/internal/pkg/db"
	"backend-go/internal/pkg/logger"
)

func main() {
	config.InitConf()
	logger.InitLogger()
	db.InitDB()

	app.InitRouter()
}
