package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type News struct {
	Id        bson.ObjectId `json:"_id"       bson:"_id"`
	Title     string        `json:"title"     bson:"title"`
	Content   string        `json:"content"   bson:"content"`
	Picture   string        `json:"picture"   bson:"picture"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	AuthorID  string        `json:"authorID"  bson:"authorID"`
}
