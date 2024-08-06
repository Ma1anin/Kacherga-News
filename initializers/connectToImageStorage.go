package initializers

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cloud *cloudinary.Cloudinary
var err error

func NewFileName() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(time.Now().String())))
}

func ConnectToImageStorage() {
	cloud, err = cloudinary.New()
	if err != nil {
		panic(err)
	}
}

func UploadImageToStore(f *multipart.FileHeader) (string, error) {
	file, err := f.Open()
	if err != nil {
		return "", errors.New("Failed to open file: " + err.Error())
	}
	defer file.Close()

	fileName := NewFileName()
	
	resp, err := cloud.Upload.Upload(context.TODO(), file, uploader.UploadParams{
        PublicID:       fileName,
        UniqueFilename: api.Bool(true),
	})
    if err != nil {
        return "", errors.New("Failed to upload image: " + err.Error())
    }

	return resp.URL, nil
}