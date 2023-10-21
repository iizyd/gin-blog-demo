package main

import (
	"backend-go/internal/app"
	"backend-go/internal/pkg/config"
	"backend-go/internal/pkg/db"
)

func main() {
	config.InitConf()
	db.InitDB()

	app.InitRouter()
}
