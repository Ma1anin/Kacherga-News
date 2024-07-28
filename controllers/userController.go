package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/cortezzIP/Kacherga-News/initializers"
	"github.com/cortezzIP/Kacherga-News/models"
)

func Signup(c gin.Context) {
	var newUser models.User

	if c.Bind(&newUser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	newUser.ID = primitive.NewObjectID()
	newUser.Picture = "default.png"
	newUser.Role = "user"

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	newUser.Password = string(hash)

	collection := initializers.DB.Collection("users")

	result, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}

	c.JSON(http.StatusOK, result)
}

func Login(c gin.Context) {

}
