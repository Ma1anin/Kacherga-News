package config

import "os"

var Config struct {
	ServerAddress   string
	DBURI          string
	ImageStorageURL string
}

func LoadConfig() {
	Config.ServerAddress = os.Getenv("SERVER_ADDR")
	Config.DBURI = os.Getenv("DBURL")
	Config.ImageStorageURL = os.Getenv("CLOUDINARY_URI")
}