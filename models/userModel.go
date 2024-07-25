package models

type User struct {
	Login    string `json:"login"    bson:"login"`
	FullName string `json:"fullName" bson:"fullName"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role"     bson:"role"`
	Picture  string `json:"picture"  bson:"picture"`
}
