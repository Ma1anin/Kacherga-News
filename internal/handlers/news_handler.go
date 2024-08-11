package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/cortezzIP/Kacherga-News/internal/database"
	"github.com/cortezzIP/Kacherga-News/internal/models"
)

func GetAllNews(c *gin.Context) {
    repo := database.NewMongoNewsRepository()
    news, err := repo.GetAllNews()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, news)
}

func GetNewsByID(c *gin.Context) {
    objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is invalid"})
		return
	}
	
    repo := database.NewMongoNewsRepository()
    news, err := repo.GetNewsByID(objId)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
        return
    }

    c.JSON(http.StatusOK, news)
}

func CreateNews(c *gin.Context) {
    var news database.NewNewsData
    if err := c.ShouldBindJSON(&news); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Println(news)

	user, _ := c.Get("user")
    repo := database.NewMongoNewsRepository()
    if userPtr, ok := user.(*models.User); ok {
        err := repo.CreateNews(*userPtr, news)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating news"})
            return
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
        return
    }

    c.JSON(http.StatusCreated, news)
}

func UpdateNews(c *gin.Context) {
    objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is invalid"})
		return
	}

    var updateData bson.M
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    repo := database.NewMongoNewsRepository()
    err = repo.UpdateNewsByID(objId, &updateData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating news"})
        return
    }

    c.JSON(http.StatusOK, updateData)
}

func DeleteNews(c *gin.Context) {
    objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is invalid"})
		return
	}

    repo := database.NewMongoNewsRepository()
    err = repo.DeleteNewsByID(objId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting news"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}
