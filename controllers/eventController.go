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

func GetAllEvents(c *gin.Context) {
	collection := initializers.DB.Collection("events")
	filter := bson.D{}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find all events",
		})

		return
	}

	var result []models.Event

	if err = cursor.All(context.TODO(), &result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to decode documents",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}

func GetEventById(c *gin.Context) {
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID is invalid",
		})

		return
	}

	collection := initializers.DB.Collection("events")
	filter := bson.D{{"_id", objId}}

	var result models.Event

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find document",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateEvent(c *gin.Context) {
	var newEvent models.Event

	if c.Bind(&newEvent) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})

		return
	}

	user, _ := c.Get("user")

	newEvent.ID = primitive.NewObjectID()
	newEvent.CreatedAt = time.Now()
	newEvent.AuthorLogin = user.(models.User).Login

	collection := initializers.DB.Collection("events")

	result, err := collection.InsertOne(context.TODO(), newEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create event",
		})

		return
	}

	c.JSON(http.StatusOK, result.InsertedID)
}

func UpdateEventById(c *gin.Context) {
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

	collection := initializers.DB.Collection("events")
	filter := bson.D{{"_id", objId}}

	result, err := collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", updateData}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update event",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteEventById(c *gin.Context) {
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID is invalid",
		})

		return
	}

	collection := initializers.DB.Collection("events")
	filter := bson.D{{"_id", objId}}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete event",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}
