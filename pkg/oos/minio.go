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
	bucket := m.bucket
	cli := m.client

	if err := cli.MakeBucket(m.ctx, bucket, minio.MakeBucketOptions{Region: m.location}); err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := cli.BucketExists(m.ctx, bucket)
		if errBucketExists == nil && exists {
			// If the Bucket has been created, the output has already created the Bucket.
			log.Printf("We already own %s\n", bucket)
			p, _ := cli.GetBucketPolicy(m.ctx, "accounts")
			log.Println("权限：" + p)
			return nil
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucket)
	}

	// Set bucket permissions.
	//// Create policy
	//policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::my-bucketname/*"],"Sid": ""}]}`

	//if err := cli.SetBucketPolicy(context.Background(), bucket, policy); err != nil {
	//	return err
	//}
	//log.Println("设置了权限")

	p, _ := cli.GetBucketPolicy(m.ctx, bucket)
	log.Println("权限：" + p)
	// 给创建的桶设置读写权限 如果不设置读写权限返回的url 将无法直接下载或在html中显示
	// json 数据由GetBucketPolicy读取而来，只需要替换 cfg.MinioBucket 即可
	//policy := `{Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:ListBucketMultipartUploads","s3:GetBucketLocation","s3:ListBucket"],"Resource":["arn:aws:s3:::"+cfg.MinioBucket+""]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:AbortMultipartUpload","s3:DeleteObject","s3:GetObject","s3:ListMultipartUploadParts","s3:PutObject"],"Resource":["arn:aws:s3:::"` + bucket +`"/*"]}]}`
	//if err := cli.SetBucketPolicy(m.ctx, bucket, policy); err!=nil {
	//	log.Printf("Bucket: %s 权限设置失败\n", bucket)
	//	log.Println(err.Error())
	//}

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
