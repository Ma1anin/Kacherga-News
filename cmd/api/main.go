package main

import (
	"github.com/cortezzIP/Kacherga-News/internal/config"
	"github.com/cortezzIP/Kacherga-News/internal/database"
	"github.com/cortezzIP/Kacherga-News/internal/initializers"
	"github.com/cortezzIP/Kacherga-News/internal/routes"
)

func init() {
	initializers.LoadEnvVariables()
	config.LoadConfig()
	database.ConnectToDb()
	database.ConnectToImageStorage()
}

func main() {
	defer database.CloseConnection()

	router := routes.SetupRoutes()
	router.Run(config.Config.ServerAddress)
}
