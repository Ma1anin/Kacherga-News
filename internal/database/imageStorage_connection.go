package database

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cortezzIP/Kacherga-News/internal/config"
)

var ImageStorageClient *cloudinary.Cloudinary

func ConnectToImageStorage() {
	var err error

	ImageStorageClient, err = cloudinary.NewFromURL(config.Config.ImageStorageURL)
	if err != nil {
		panic(err)
	}
}
