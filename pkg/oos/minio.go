package oos

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"log"
	"mime/multipart"
)

type min struct {
	ctx      context.Context
	client   *minio.Client
	file     *multipart.FileHeader
	bucket   string
	location string
	fileType string
}

func (m *min) MakeBucket() error {
	cli := m.client

	if err := cli.MakeBucket(m.ctx, m.bucket, minio.MakeBucketOptions{Region: m.location}); err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := cli.BucketExists(m.ctx, m.bucket)
		if errBucketExists == nil && exists {
			// If the Bucket has been created, the output has already created the Bucket.
			log.Printf("We already own %s\n", m.bucket)
			return nil
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", m.bucket)
		// Set bucket permissions.
		// Set the access permissions of the bucket.
		// By default, it has read and write permissions so that the returned url can be
		// accessed and downloaded in the browser. It can be changed according to business needs.
		policy := `{"Version":"2012-10-17",
"Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},
"Action":["s3:ListBucket","s3:ListBucketMultipartUploads","s3:GetBucketLocation"],
"Resource":["arn:aws:s3:::accounts"]},{"Effect":"Allow",
"Principal":{"AWS":["*"]},
"Action":["s3:AbortMultipartUpload","s3:DeleteObject","s3:GetObject","s3:ListMultipartUploadParts","s3:PutObject"],
"Resource":["arn:aws:s3:::accounts/*"]}]}
`

		if err := cli.SetBucketPolicy(context.Background(), m.bucket, policy); err != nil {
			return err
		}
		log.Printf("Successfully set permissions %s\n", m.bucket)

		// The method used to obtain minio permissions,
		// returns the detailed permissions of the bucket,
		// which can be used to obtain settings,
		// in development and used in the method ( SetBucketPolicy() ) of setting permissions.
		// p, err := client.GetBucketPolicy(context.Background(), m.bucket)
	}
	return nil
}

func (m *min) FileUploader() (string, error) {
	err := m.MakeBucket()
	if err != nil {
		log.Println("Error in MakeBucket:", err)
		return "", err
	}

	file, err := m.file.Open()
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer file.Close()

	// file Name
	fn := fmt.Sprintf("%s/%s-%s", m.fileType, uuid.New().String(), m.file.Filename)

	info, err := m.client.PutObject(m.ctx, m.bucket, fn, file, -1, minio.PutObjectOptions{ContentType: m.fileType})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	log.Printf("Successfully uploaded %s of size %d\n", m.file.Filename, info.Size)
	return fn, nil
}

type Min interface {
	// MakeBucket Make a new bucket
	MakeBucket() error
	// FileUploader // Upload file
	FileUploader() (string, error)
}

func NewMin(file *multipart.FileHeader, fileType string, bucket string, location string) Min {
	ctx := context.Background()
	client, err := initMinio()
	if err != nil {
		return nil
	}
	return &min{ctx: ctx, client: client, file: file, fileType: fileType, bucket: bucket, location: location}
}

// initMinio Initialize minio and return to client.
func initMinio() (*minio.Client, error) {
	endpoint := viper.GetString("oos.minio.addr")
	accessKeyID := viper.GetString("oos.minio.accessKeyID")
	secretAccessKey := viper.GetString("oos.minio.secretAccessKey")
	useSSL := viper.GetBool("oos.minio.useSSL")

	cli, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return nil, err
	}
	return cli, nil
}
