package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	Login    string
	ImageURL string
}

type News struct {
	ID        primitive.ObjectID `json:"_id"       bson:"_id"`
	Title     string             `json:"title"     bson:"title"`
	Content   string             `json:"content"   bson:"content"`
	ImageURL  string             `json:"picture"   bson:"picture"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	Author    Author             `json:"author"    bson:"author"`
}
