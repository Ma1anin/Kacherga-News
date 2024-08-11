package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/cortezzIP/Kacherga-News/internal/database"
	"github.com/cortezzIP/Kacherga-News/internal/models"
)

func GetAllEvents(c *gin.Context) {
    repo := database.NewMongoEventRepository()
    news, err := repo.GetAllEvents()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, news)
}

func GetEventByID(c *gin.Context) {
    objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is invalid"})
		return
	}
	
    repo := database.NewMongoEventRepository()
    event, err := repo.GetEventByID(objId)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
        return
    }

    c.JSON(http.StatusOK, event)
}

func CreateEvent(c *gin.Context) {
    var event database.NewEventData
    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	user, _ := c.Get("user")
    repo := database.NewMongoEventRepository()
    if userPtr, ok := user.(*models.User); ok {
        err := repo.CreateEvent(*userPtr, event)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating event"})
            return
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
        return
    }

    c.JSON(http.StatusCreated, event)
}

func UpdateEvent(c *gin.Context) {
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

    repo := database.NewMongoEventRepository()
    err = repo.UpdateEventByID(objId, &updateData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating event"})
        return
    }

    c.JSON(http.StatusOK, updateData)
}

func DeleteEvent(c *gin.Context) {
    objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is invalid"})
		return
	}

    repo := database.NewMongoEventRepository()
    err = repo.DeleteEventByID(objId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting event"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
