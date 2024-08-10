package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/cortezzIP/Kacherga-News/internal/models"
)

type UserRepository interface {
	CreateUser(newUserData NewUserData) error
	UpdateUser(id primitive.ObjectID, updateData bson.M) error
	GetUserByLogin(login string) (*models.User, error)
	DeleteUserByID(id primitive.ObjectID) error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

type NewUserData struct {
	Login    string
	FullName string
	Password string
}

func NewMongoUserRepository() *MongoUserRepository {
	return &MongoUserRepository{
		collection: MongoClient.Database("Cluster0").Collection("users"),
	}
}

func (repo MongoUserRepository) GetUserByLogin(login string) (*models.User, error) {
	var result *models.User
	filter := bson.D{{Key: "login", Value: login}}

	if err := repo.collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, errors.New("Failed to decode user: " + err.Error())
	}

	return result, nil
}

func (repo MongoUserRepository) CreateUser(newUserData NewUserData) error {
	var findResult models.User
	filter := bson.D{{Key: "login", Value: newUserData.Login}}

	repo.collection.FindOne(context.TODO(), filter).Decode(&findResult)
	if findResult.ID != primitive.NilObjectID {
		return errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUserData.Password), 10)
	if err != nil {
		return errors.New("failed to hash password")
	}

	var newUser models.User

	newUser.ID = primitive.NewObjectID()
	newUser.Login = newUserData.Login
	newUser.FullName = newUserData.FullName
	newUser.Password = string(hash)
	newUser.Role = "user"
	newUser.ImageURL = "https://res.cloudinary.com/ddmrmjewm/image/upload/v1723101640/default.png"

	_, err = repo.collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return errors.New("failed to create user: " + err.Error())
	}

	return nil
}

func (repo MongoUserRepository) UpdateUser(id primitive.ObjectID, updateData bson.M) error {
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repo.collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: updateData}})
	if err != nil {
		return errors.New("Failed to update user: " + err.Error())
	}

	return nil
}

func (repo MongoUserRepository) DeleteUserByID(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repo.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("Failed to delete user: " + err.Error())
	}

	return nil
}
