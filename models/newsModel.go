package models

import (
	"time"
)

type News struct {
	Title     string    `json:"title"     bson:"title"`
	Content   string    `json:"content"   bson:"content"`
	Picture   string    `json:"picture"   bson:"picture"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	AuthorID  string    `json:"authorID"  bson:"authorID"`
}
