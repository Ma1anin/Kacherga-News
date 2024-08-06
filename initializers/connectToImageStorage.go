package initializers

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
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

func UploadImageToStore(base64String string) (string, error) {
	fileName := NewFileName()
	
	resp, err := cloud.Upload.Upload(context.TODO(), base64String, uploader.UploadParams{
        PublicID:       fileName,
        UniqueFilename: api.Bool(true),
	})
    if err != nil {
        return "", errors.New("Failed to upload image: " + err.Error())
    }

	return resp.URL, nil
}