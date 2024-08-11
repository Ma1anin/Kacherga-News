package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/cortezzIP/Kacherga-News/internal/database"
	"github.com/cortezzIP/Kacherga-News/internal/models"
)

const cookieMaxAge = 3600 * 24 * 3

func Signup(c *gin.Context) {
	var user database.NewUserData
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	repo := database.NewMongoUserRepository()
	err := repo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
        return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var requestBody struct {
		Login    string
		Password string
	}

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	repo := database.NewMongoUserRepository()
	user, err := repo.GetUserByLogin(requestBody.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login or password is incorrect"})
		return
	}
	
	if user.ID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login or password is incorrect"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login or password is incorrect"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	user, _ := c.Get("user")

	repo := database.NewMongoUserRepository()
	if userPtr, ok := user.(*models.User); ok {
        err := repo.UpdateUser(userPtr.ID, updateData)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error update user"})
            return
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
        return
    }

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func DeleteUserByID(c *gin.Context) {
	user, _ := c.Get("user")

	repo := database.NewMongoUserRepository()
	if userPtr, ok := user.(*models.User); ok {
        err := repo.DeleteUserByID(userPtr.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error delete user"})
            return
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
        return
    }

	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
}
