package routes

import (
	"github.com/cortezzIP/Kacherga-News/internal/handlers"
	"github.com/cortezzIP/Kacherga-News/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"}
	r.Use(cors.New(config))

	r.MaxMultipartMemory = 8 << 20

	router := r.Group("/api")

	router.GET("/news", handlers.GetAllNews)
	router.GET("/news/:id", handlers.GetNewsByID)
	router.POST("/news", middleware.RequireAuth, handlers.CreateNews)
	router.PUT("/news/:id", middleware.RequireAuth, handlers.UpdateNews)
	router.DELETE("/news/:id", middleware.RequireAuth, handlers.DeleteNews)

	router.GET("/event", handlers.GetAllEvents)
	router.GET("/event/:id", handlers.GetEventByID)
	router.POST("/event", middleware.RequireAuth, handlers.CreateEvent)
	router.PUT("/event/:id", middleware.RequireAuth, handlers.UpdateEvent)
	router.DELETE("/event/:id", middleware.RequireAuth, handlers.DeleteEvent)

	router.GET("/validate", middleware.RequireAuth, handlers.Validate)
	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)
	router.POST("/logout", middleware.RequireAuth, handlers.Logout)
	router.PUT("/updateAccount", middleware.RequireAuth, handlers.UpdateUser)
	router.DELETE("/deleteAccount", middleware.RequireAuth, handlers.DeleteUserByID)

	return r
}