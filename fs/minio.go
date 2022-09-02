package fs

import (
	"fmt"
	"github.com/hvxahv/hvx/errors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"io"
)

const (
	MinioPublicAvatarBucketPolicy = `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetBucketLocation","s3:ListBucket","s3:ListBucketMultipartUploads"],"Resource":["arn:aws:s3:::avatar"]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:AbortMultipartUpload","s3:DeleteObject","s3:GetObject","s3:ListMultipartUploadParts","s3:PutObject"],"Resource":["arn:aws:s3:::avatar/*"]}]}`
	MinioPublicAttachBucketPolicy = `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetBucketLocation","s3:ListBucket","s3:ListBucketMultipartUploads"],"Resource":["arn:aws:s3:::attach"]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:AbortMultipartUpload","s3:DeleteObject","s3:GetObject","s3:ListMultipartUploadParts","s3:PutObject"],"Resource":["arn:aws:s3:::attach/*"]}]}`
)

type Minio struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
	Region          string
	Ctx             context.Context
	Client          *minio.Client
}

type MinioBucket interface {
	MakeBucket(name, policy string) error
	ListBuckets() error
	GetBucketPolicy(name string) (string, error)
}

type MinioFiles struct {
	Ctx         context.Context
	Client      *minio.Client
	BucketName  string
	ObjectName  string
	ContentType string

	Buffer io.Reader
	Size   int64
}

type MinioFile interface {
	Put() (*minio.UploadInfo, error)
	Remove(o string)
}

func NewDefaultMinio() *Minio {
	var (
		endpoint        = viper.GetString("fs.minio.endpoint")
		accessKeyID     = viper.GetString("fs.minio.accessKeyID")
		secretAccessKey = viper.GetString("fs.minio.secretAccessKey")
		useSSL          = viper.GetBool("fs.minio.useSSL")
		region          = viper.GetString("fs.minio.region")
	)
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
		return nil, errors.Newf("minio connection error: %v ", err)
	}
	x.Client = client
	x.Ctx = context.Background()
	return x, err
}

func (x *Minio) MakeBucket(name, policy string) error {
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
	if err := x.Client.SetBucketPolicy(x.Ctx, name, policy); err != nil {
		return err
	}
	return nil
}

func (x *Minio) ListBuckets() error {
	buckets, err := x.Client.ListBuckets(x.Ctx)
	if err != nil {
		return err
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
	return nil
}

func (x *Minio) GetBucketPolicy(name string) (string, error) {
	policy, err := x.Client.GetBucketPolicy(x.Ctx, name)
	if err != nil {
		return "", err
	}
	return policy, nil
}

func NewMinioFiles(client *minio.Client, ctx context.Context, bucketName, objectName, contentType string, buffer io.Reader, size int64) *MinioFiles {
	return &MinioFiles{Client: client, Ctx: ctx, BucketName: bucketName, ObjectName: objectName, ContentType: contentType, Buffer: buffer, Size: size}
}

func (x *MinioFiles) Put() (*minio.UploadInfo, error) {
	o, err := x.Client.PutObject(x.Ctx, x.BucketName, x.ObjectName, x.Buffer, x.Size, minio.PutObjectOptions{
		ContentType: x.ContentType,
	})
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func NewMinioRemoveFile(client *minio.Client, ctx context.Context, bucket, name string) *MinioFiles {
	return &MinioFiles{
		Ctx:        ctx,
		Client:     client,
		BucketName: bucket,
		ObjectName: name,
	}
}

func (x *MinioFiles) Remove() error {
	if err := x.Client.RemoveObject(x.Ctx, x.BucketName, x.ObjectName, minio.RemoveObjectOptions{
		ForceDelete: true}); err != nil {
		return err
	}

	return nil
}
