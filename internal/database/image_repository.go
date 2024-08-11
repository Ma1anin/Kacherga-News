package database

import (
	"context"
	"errors"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/cortezzIP/Kacherga-News/internal/utils"
)

type ImageRepository interface {
	UploadImageToStorage(string) (string, error)
}

type CloudinaryImageRepository struct {
	ImageStorageClient *cloudinary.Cloudinary
}

func NewImageStorage() *CloudinaryImageRepository {
	return &CloudinaryImageRepository{
		ImageStorageClient: ImageStorageClient,
	}
}

func (repo CloudinaryImageRepository) UploadImageToStorage(base64String string) (string, error) {
	fileName := utils.NewFileName()

	resp, err := repo.ImageStorageClient.Upload.Upload(context.TODO(), base64String, uploader.UploadParams{
		PublicID:       fileName,
		UniqueFilename: api.Bool(true),
	})
	if err != nil {
		return "", errors.New("Failed to upload image: " + err.Error())
	}

	return resp.URL, nil
}
