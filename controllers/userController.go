package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/cortezzIP/Kacherga-News/initializers"
	"github.com/cortezzIP/Kacherga-News/models"
)

const cookieMaxAge = 3600 * 24 * 3

func Signup(c *gin.Context) {
	var newUser models.User

	if c.Bind(&newUser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var findResult models.User
	collection := initializers.DB.Collection("users")
	filter := bson.D{{Key: "login", Value: newUser.Login}}

	collection.FindOne(context.TODO(), filter).Decode(&findResult)
	if findResult.ID != primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already register",
		})

		return
	}

	newUser.ID = primitive.NewObjectID()
	newUser.ImageURL = "https://res.cloudinary.com/ddmrmjewm/image/upload/v1723101640/default.png"
	newUser.Role = "user"

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	newUser.Password = string(hash)

	result, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}

	c.JSON(http.StatusOK, result.InsertedID)
}

func Login(c *gin.Context) {
	var requestBody struct {
		Login    string
		Password string
	}

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})

		return
	}

	var user models.User
	collection := initializers.DB.Collection("users")
	filter := bson.D{{Key: "login", Value: requestBody.Login}}

	collection.FindOne(context.TODO(), filter).Decode(&user)

	if user.ID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Login or password is incorrect",
		})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Login or password is incorrect",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Login,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, cookieMaxAge, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func UpdateUser(c *gin.Context) {
	var updateData bson.M
	if c.ShouldBindJSON(&updateData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})

		return
	}

	user, _ := c.Get("user")
	collection := initializers.DB.Collection("users")
	filter := bson.D{{Key: "_id", Value: user.(models.User).ID}}

	_, err := collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: updateData}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update user",
		})

		return
	}

	updateNewsQuery := bson.D{{Key: "$set", Value: bson.D{{Key: "author.imageURL", Value: updateData["imageURL"]}}}}
    filterForNews := bson.D{{Key: "author.login", Value: user.(models.User).Login}}
    _, err = collection.UpdateMany(context.TODO(), filterForNews, updateNewsQuery)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to update related news",
        })
        return
    }

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
}
