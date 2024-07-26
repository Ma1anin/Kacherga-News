package main

import (
	"os"

	"github.com/cortezzIP/Kacherga-News/controllers"
	"github.com/cortezzIP/Kacherga-News/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/news", controllers.CreateNews)
	r.GET("/news", controllers.GetAllNews)
	r.GET("/news/:id", controllers.GetNewsById)
	r.PUT("/news/:id", controllers.UpdateNewsById)
	r.DELETE("/news/:id", controllers.DeleteNewsById)

	r.Run(os.Getenv("LISTEN_ADDR"))
}
