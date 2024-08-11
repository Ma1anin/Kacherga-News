package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID          primitive.ObjectID `json:"_id"          bson:"_id"`
	Title       string             `json:"title"        bson:"title"`
	Content     string             `json:"content"      bson:"content"`
	CreatedAt   string             `json:"createdAt"    bson:"createdAt"`
	AuthorLogin string             `json:"authorLogin"  bson:"authorLogin"`
}