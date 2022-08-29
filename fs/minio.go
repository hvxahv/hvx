package fs

import (
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

const (
	defaultEndpoint        = "minio.halfmemories.com"
	defaultAccessKeyID     = "hvxahv"
	defaultSecretAccessKey = "hvxahv123"
)

type Minio struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
	Region          string

	Ctx    context.Context
	Client *minio.Client

	BucketName  string
	ObjectName  string
	ContentType string
}

func (x *Minio) SetPutOption(bucketName, objectName, contentType string) *Minio {
	x.ObjectName = objectName
	x.BucketName = bucketName
	x.ContentType = contentType
	return x
}

func NewMinio() *Minio {
	var (
		endpoint        = viper.GetString("minio.endpoint")
		accessKeyID     = viper.GetString("minio.accessKeyID")
		secretAccessKey = viper.GetString("minio.secretAccessKey")
		useSSL          = viper.GetBool("minio.useSSL")
		region          = viper.GetString("minio.region")
	)

	if endpoint == "" {
		endpoint = defaultEndpoint
	}
	if accessKeyID == "" {
		accessKeyID = defaultAccessKeyID
	}
	if secretAccessKey == "" {
		secretAccessKey = defaultSecretAccessKey
	}
	return &Minio{
		endpoint:        endpoint,
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		useSSL:          useSSL,
		Region:          region,
	}
}

// Dial the server connected to Minio returns the client with the context.
func (x *Minio) Dial() (*Minio, error) {
	client, err := minio.New(x.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(x.accessKeyID, x.secretAccessKey, ""),
		Secure: x.useSSL})
	if err != nil {
		return nil, errors.Errorf("minio connection error: %v ", err)
	}
	x.Client = client
	x.Ctx = context.Background()
	return x, err
}

type Miner interface {
	CreateBucket(name string) error
	ListBucket() error
}

func (x *Minio) CreateBucket(name string) error {
	if err := x.Client.MakeBucket(x.Ctx, name, minio.MakeBucketOptions{
		Region:        x.Region,
		ObjectLocking: false,
	}); err != nil {
		exists, _ := x.Client.BucketExists(x.Ctx, name)
		if exists {
			return errors.New("BUCKET_ALREADY_EXIST")
		}
		return err
	}
	return nil
}

func (x *Minio) ListBucket() error {
	buckets, err := x.Client.ListBuckets(x.Ctx)
	if err != nil {
		return err
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
	return nil
}

func (x *Minio) FilePut(buffer io.Reader, size int64) (*minio.UploadInfo, error) {
	o, err := x.Client.PutObject(x.Ctx, x.BucketName, x.ObjectName, buffer, size, minio.PutObjectOptions{
		ContentType: x.ContentType,
	})

	if err != nil {
		return nil, err
	}
	return &o, nil
}
