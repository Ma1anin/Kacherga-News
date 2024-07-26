package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type News struct {
	ID        primitive.ObjectID `json:"_id"       bson:"_id"`
	Title     string             `json:"title"     bson:"title"`
	Content   string             `json:"content"   bson:"content"`
	Picture   string             `json:"picture"   bson:"picture"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	AuthorID  string             `json:"authorID"  bson:"authorID"`
}
