package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"log"
	"mime/multipart"
)

type Minio struct {
	Ctx      context.Context
	Client   *minio.Client
	File     *multipart.FileHeader
	Bucket   string
	Location string
	FileType string
}

func NewMinio(file *multipart.FileHeader, bucket string, location string, fileType string) *Minio {
	ctx := context.Background()
	client, err := InitMinio()
	if err != nil {
		log.Println(err)
	}
	return &Minio{Ctx: ctx, Client: client, File: file, Bucket: bucket, Location: location, FileType: fileType}
}

type I interface {
	// NewBucket Make a new bucket.
	NewBucket() error
	// Uploader // Upload file and return address.
	Uploader() (string, error)
}

// InitMinio Initialize minio and return to client.
func InitMinio() (*minio.Client, error) {
	endpoint := viper.GetString("minio.addr")
	accessKeyID := viper.GetString("minio.accessKeyID")
	secretAccessKey := viper.GetString("minio.secretAccessKey")
	useSSL := viper.GetBool("minio.useSSL")

	cli, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		_ = fmt.Errorf("failed to connect to minio: %v", err)
		return nil, err
	}
	return cli, nil
}
