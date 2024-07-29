package main

import (
	"os"

	"github.com/cortezzIP/Kacherga-News/controllers"
	"github.com/cortezzIP/Kacherga-News/initializers"
	"github.com/cortezzIP/Kacherga-News/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/news", controllers.GetAllNews)
	r.GET("/news/:id", controllers.GetNewsById)
	r.POST("/news", controllers.CreateNews)
	r.PUT("/news/:id", controllers.UpdateNewsById)
	r.DELETE("/news/:id", controllers.DeleteNewsById)

	r.GET("/event", controllers.GetAllEvents)
	r.GET("/event/:id", controllers.GetEventById)
	r.POST("/event", controllers.CreateEvent)
	r.PUT("/event/:id", controllers.UpdateEventById)
	r.DELETE("/event/:id", controllers.DeleteEventById)
	
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	
	r.Run(os.Getenv("LISTEN_ADDR"))
}
