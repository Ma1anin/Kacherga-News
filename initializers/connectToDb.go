package initializers

import (
	"os"

	"gopkg.in/mgo.v2"
)

var DB *mgo.Database

func ConnectToDb() {
	session, err := mgo.Dial(os.Getenv("DB_URL"))

	if err != nil {
		panic(err)
	}
	defer session.Close()

	DB = session.DB("Cluster0")
}