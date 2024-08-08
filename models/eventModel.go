package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID `json:"_id"       bson:"_id"`
	Title       string             `json:"title"     bson:"title"`
	Content     string             `json:"content"   bson:"content"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	AuthorLogin string             `json:"authorID"  bson:"authorID"`
}
