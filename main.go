package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/cortezzIP/Kacherga-News/controllers"
	"github.com/cortezzIP/Kacherga-News/initializers"
	"github.com/cortezzIP/Kacherga-News/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.ConnectToImageStorage()
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://127.0.0.1:5500"}
    r.Use(cors.New(config))

	r.MaxMultipartMemory = 8 << 20

	router := r.Group("/api")

	router.GET("/news", controllers.GetAllNews)
	router.GET("/news/:id", controllers.GetNewsById)
	router.POST("/news", middleware.RequireAuth, controllers.CreateNews)
	router.PUT("/news/:id", middleware.RequireAuth, controllers.UpdateNewsById)
	router.DELETE("/news/:id", middleware.RequireAuth, controllers.DeleteNewsById)

	router.GET("/event", controllers.GetAllEvents)
	router.GET("/event/:id", controllers.GetEventById)
	router.POST("/event", middleware.RequireAuth, controllers.CreateEvent)
	router.PUT("/event/:id", middleware.RequireAuth, controllers.UpdateEventById)
	router.DELETE("/event/:id", middleware.RequireAuth, controllers.DeleteEventById)
	
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.POST("/logout", middleware.RequireAuth, controllers.Logout)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	
	r.Run(os.Getenv("LISTEN_ADDR"))
}
