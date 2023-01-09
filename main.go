package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	initModel "github.com/wibowo-id/go_core/app/models"
	cfg "github.com/wibowo-id/go_core/config"
	"github.com/wibowo-id/go_core/routes"
)

type App struct {
	config cfg.Config
}

var app App

func init() {
	config := cfg.NewViperConfig()
	app = App{config: config}

	if config.GetBool(`app.debug`) {
		fmt.Println("Service RUN on DEBUG mode. Port: " + config.GetString("app.host"))
	}
}

func main() {
	dbPostgreSQL := initModel.PgInfo{
		Hostname: app.config.GetString("database.postgresql.host"),
		Username: app.config.GetString("database.postgresql.username"),
		Password: app.config.GetString("database.postgresql.password"),
		Database: app.config.GetString("database.postgresql.database"),
		Port:     app.config.GetString("database.postgresql.port"),
	}
	dbPostgreSQL.ConnectDatabase()

	r := gin.Default()
	r.Use(cors.Default())
	routes.ApiRoutes(r)
	err := r.Run(app.config.GetString("app.host"))
	if err != nil {
		return
	}
}
