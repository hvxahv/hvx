package fs

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Minio struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
}

func NewMinio() *Minio {
	var (
		endpoint        = viper.GetString("minio.endpoint")
		accessKeyID     = viper.GetString("minio.accessKeyID")
		secretAccessKey = viper.GetString("minio.secretAccessKey")
		useSSL          = viper.GetBool("minio.useSSL")
	)
	fmt.Println(endpoint, accessKeyID, secretAccessKey, useSSL)
	return &Minio{endpoint: endpoint, accessKeyID: accessKeyID, secretAccessKey: secretAccessKey, useSSL: useSSL}
}

var minioClient *minio.Client

func (x *Minio) Dial() error {
	cli, err := minio.New(x.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(x.accessKeyID, x.secretAccessKey, ""),
		Secure: x.useSSL,
	})
	if err != nil {
		return errors.Errorf("minio connection error: %v ", err)
	}
	minioClient = cli
	return nil
}

func GetMinio() *minio.Client {
	return minioClient
}
