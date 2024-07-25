package main

import (
	"os"

	"github.com/cortezzIP/Kacherga-News/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.Run(os.Getenv("PORT"))
}
