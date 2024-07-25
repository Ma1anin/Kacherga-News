package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `json:"_id"      bson:"_id"`
	Login    string        `json:"login"    bson:"login"`
	FullName string        `json:"fullName" bson:"fullName"`
	Password string        `json:"password" bson:"password"`
	Role     string        `json:"role"     bson:"role"`
	Picture  string        `json:"picture"  bson:"picture"`
}
