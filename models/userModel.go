package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Login    string             `json:"login"    bson:"login"`
	FullName string             `json:"fullName" bson:"fullName"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role"     bson:"role"`
	Picture  string             `json:"picture"  bson:"picture"`
}
