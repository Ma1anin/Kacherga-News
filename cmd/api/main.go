package main

import (
	"github.com/cortezzIP/Kacherga-News/internal/config"
	"github.com/cortezzIP/Kacherga-News/internal/routes"
	"github.com/cortezzIP/Kacherga-News/internal/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := routes.SetupRoutes()
	router.Run(config.Config.ServerAddress)
}
