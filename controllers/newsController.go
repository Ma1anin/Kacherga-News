package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/cortezzIP/Kacherga-News/initializers"
	"github.com/cortezzIP/Kacherga-News/models"
)

func GetAllNews(c *gin.Context) {
	collection := initializers.DB.Collection("news")
	filter := bson.D{}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find all news",
		})

		return
	}

	var result []models.News
	if err = cursor.All(context.TODO(), &result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to decode documents",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}

func GetNewsById(c *gin.Context) {
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID is invalid",
		})

		return
	}

	collection := initializers.DB.Collection("news")
	filter := bson.D{{"_id", objId}}

	var result models.News

	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find news",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateNews(c *gin.Context) {
	var newNews models.News

	if c.Bind(&newNews) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})

		return
	}
	
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get file from form: " + err.Error(),
		})
		return
	}

	imageURL, err := initializers.UploadImageToStore(file)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add file to storage: " + err.Error(),
		})
		return
	}
	
	collection := initializers.DB.Collection("news")
	user, _ := c.Get("user")

	newNews.ID = primitive.NewObjectID()
	newNews.Author.Login = user.(models.User).Login
	newNews.Author.ImageURL = user.(models.User).ImageURL
	newNews.ImageURL = imageURL
	newNews.CreatedAt = time.Now()

	result, err := collection.InsertOne(context.TODO(), newNews)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create news",
		})

		return
	}

	c.JSON(http.StatusOK, result.InsertedID)
}

func UpdateNewsById(c *gin.Context) {
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID is invalid",
		})

		return
	}

	var updateData bson.M
	if c.ShouldBindJSON(&updateData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})

		return
	}

	collection := initializers.DB.Collection("news")
	filter := bson.D{{"_id", objId}}

	result, err := collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", updateData}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update news",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteNewsById(c *gin.Context) {
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID is invalid",
		})

		return
	}

	collection := initializers.DB.Collection("news")
	filter := bson.D{{"_id", objId}}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete news",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}
