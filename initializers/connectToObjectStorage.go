package initializers

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const bucketName string = "kacherga-news-bucket"

var ObjectStorageUploader *manager.Uploader
var ObjectStorageDownloader *manager.Downloader

func ConnectToObjectStorage() {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "ru-central-1" {
			return aws.Endpoint{
				PartitionID:   "cloud",
				URL:           "https://s3.cloud.ru/",
				SigningRegion: "ru-central-1",
			}, nil
		}
		return aws.Endpoint{}, errors.New("unknown endpoint requested")
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)

	ObjectStorageUploader = manager.NewUploader(client)
	ObjectStorageDownloader = manager.NewDownloader(client)
}

func AddFileToStorage(f *multipart.FileHeader) (string, error) {
	file, openErr := f.Open()
	if openErr != nil {
		return "", errors.New("File error: " + openErr.Error())
	}

	defer file.Close()

	result, uploadErr := ObjectStorageUploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fmt.Sprintf("%x", sha256.Sum256([]byte(time.Now().String())))),
		Body:   file,
		ACL:    "public-read",
	})

	if uploadErr != nil {
		return "", errors.New("Upload error: " + uploadErr.Error())
	}

	return *result.Key, nil
}

func DownloadFileFromStorage(key *string) {

}
