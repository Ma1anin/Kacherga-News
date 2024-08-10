package database

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/cortezzIP/Kacherga-News/internal/models"
)

type EventRepository interface {
	GetAllEvents() ([]*models.Event, error)
	GetEventByID(id primitive.ObjectID) (*models.Event, error)
	CreateEvent(user models.User, newEventData NewEventData) error
	UpdateEventByID(id primitive.ObjectID, updateData *bson.M) error
	DeleteEventByID(id primitive.ObjectID) error
}

type MongoEventRepository struct {
	collection *mongo.Collection
}

type NewEventData struct {
	Title   string
	Content string
}

func NewMongoEventRepository() *MongoEventRepository {
	return &MongoEventRepository{
		collection: MongoClient.Database("Cluster0").Collection("events"),
	}
}

func (repo MongoEventRepository) GetAllEvents() ([]*models.Event, error) {
	cursor, err := repo.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, errors.New("Filed to get events: " + err.Error())
	}

	var results []*models.Event
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, errors.New("Filed to decode events: " + err.Error())
	}

	return results, nil
}

func (repo MongoEventRepository) GetEventByID(id primitive.ObjectID) (*models.Event, error) {
	var result *models.Event
	filter := bson.D{{Key: "_id", Value: id}}

	err := repo.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, errors.New("Filed to get event: " + err.Error())
	}

	return result, nil
}

func (repo MongoEventRepository) CreateEvent(user models.User, newEventData NewEventData) error {
	var newEvent models.Event

	newEvent.ID = primitive.NewObjectID()
	newEvent.AuthorLogin = user.Login
	newEvent.Title = newEventData.Title
	newEvent.Content = newEventData.Content
	newEvent.CreatedAt = time.Now()

	_, err := repo.collection.InsertOne(context.TODO(), newEvent)
	if err != nil {
		return errors.New("Filed to create event: " + err.Error())
	}

	return nil
}

func (repo MongoEventRepository) UpdateEventByID(id primitive.ObjectID, updateData *bson.M) error {
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repo.collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: updateData}})
	if err != nil {
		return errors.New("Filed to update event: " + err.Error())
	}

	return nil
}

func (repo MongoEventRepository) DeleteEventByID(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repo.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("Filed to update event: " + err.Error())
	}

	return nil
}
