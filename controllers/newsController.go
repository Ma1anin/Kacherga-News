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
	filter := bson.D{{Key: "_id", Value: objId}}

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
	var requestBody struct {
		Title             string
		Content           string
		ImageBase64string string
	}

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})

		return
	}

	imageURL, err := initializers.UploadImageToStore(requestBody.ImageBase64string)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add file to storage: " + err.Error(),
		})
		return
	}

	collection := initializers.DB.Collection("news")
	user, _ := c.Get("user")

	newNews.ID = primitive.NewObjectID()
	newNews.Title = requestBody.Title
	newNews.Content = requestBody.Content
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
	filter := bson.D{{Key: "_id", Value: objId}}

	_, err = collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: updateData}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update news",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
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
	filter := bson.D{{Key: "_id", Value: objId}}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete news",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}
