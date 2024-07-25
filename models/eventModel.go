package models

import (
	"time"
)

type Event struct {
	Title     string    `json:"title"     bson:"title"`
	Content   string    `json:"content"   bson:"content"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	AuthorID  string    `json:"authorID"  bson:"authorID"`
}
