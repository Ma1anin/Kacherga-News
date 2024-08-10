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

type NewsRepository interface {
	GetAllNews() ([]*models.News, error)
	GetNewsByID(id primitive.ObjectID) (*models.News, error)
	CreateNews(user models.User, newNewsData NewNewsData) error
	UpdateNewsByID(id primitive.ObjectID, updateData *bson.M) error
	DeleteNewsByID(id primitive.ObjectID) error
}

type MongoNewsRepository struct {
	collection *mongo.Collection
}

type NewNewsData struct {
	Title             string
	Content           string
	ImageBase64string string
}

var cloudinaryRepo = CloudinaryImageRepository{
	ImageStorageClient: ImageStorageClient,
}

func NewMongoNewsRepository() *MongoNewsRepository {
	return &MongoNewsRepository{
		collection: MongoClient.Database("Cluster0").Collection("news"),
	}
}

func (repo MongoNewsRepository) GetAllNews() ([]*models.News, error) {
	cursor, err := repo.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, errors.New("Failed to get news: " + err.Error())
	}

	var results []*models.News
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, errors.New("Failed to decode news: " + err.Error())
	}

	return results, nil
}

func (repo MongoNewsRepository) GetNewsByID(id primitive.ObjectID) (*models.News, error) {
	var result *models.News
	filter := bson.D{{Key: "_id", Value: id}}

	if err := repo.collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, errors.New("Failed to decode news: " + err.Error())
	}

	return result, nil
}

func (repo MongoNewsRepository) CreateNews(user models.User, newNewsData NewNewsData) error {
	imageURL, err := cloudinaryRepo.UploadImageToStorage(newNewsData.ImageBase64string)
	if err != nil {
		return errors.New("Failed to add file to storage: " + err.Error())
	}

	var newNews models.News

	newNews.ID = primitive.NewObjectID()
	newNews.Title = newNewsData.Title
	newNews.Content = newNewsData.Content
	newNews.Author.Login = user.Login
	newNews.Author.ImageURL = user.ImageURL
	newNews.ImageURL = imageURL
	newNews.CreatedAt = time.Now()

	_, err = repo.collection.InsertOne(context.TODO(), newNews)
	if err != nil {
		return errors.New("Failed to create news: " + err.Error())
	}

	return nil
}

func (repo MongoNewsRepository) UpdateNewsByID(id primitive.ObjectID, updateData *bson.M) error {
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repo.collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: updateData}})
	if err != nil {
		return errors.New("Failed to update news: " + err.Error())
	}

	return nil
}

func (repo MongoNewsRepository) DeleteNewsByID(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repo.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("Failed to delete news: " + err.Error())
	}

	return nil
}
