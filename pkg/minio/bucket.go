package minio

import (
	"github.com/minio/minio-go/v7"
	"golang.org/x/net/context"
	"log"
)

type Buckets struct {
	Context  context.Context
	Client   *minio.Client
	Policy   string
	Bucket   string
	Location string
}

func NewBucket(bucket, location string) *Buckets {
	client, err := InitMinIO()
	if err != nil {
		log.Println(err)
		return nil
	}

	policy := `{"Version":"2012-10-17",
"Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},
"Action":["s3:GetBucketLocation","s3:ListBucket","s3:ListBucketMultipartUploads"],
"Resource":["arn:aws:s3:::avatar"]},{"Effect":"Allow","Principal":{"AWS":["*"]},
"Action":["s3:AbortMultipartUpload","s3:DeleteObject","s3:GetObject","s3:ListMultipartUploadParts","s3:PutObject"],
"Resource":["arn:aws:s3:::avatar/*"]}]}
`
	return &Buckets{Context: context.Background(), Client: client, Policy: policy, Bucket: bucket, Location: location}
}

type Bucket interface {
	// Create Make a new bucket.
	Create() error
}

func (m *Buckets) Create() error {
	cli := m.Client

	if err := cli.MakeBucket(m.Context, m.Bucket, minio.MakeBucketOptions{Region: m.Location}); err != nil {
		exists, errBucketExists := cli.BucketExists(m.Context, m.Bucket)
		if errBucketExists == nil && exists {
			return nil
		}
		return err
	}
	policy := m.Policy
	if err := cli.SetBucketPolicy(context.Background(), m.Bucket, policy); err != nil {
		return err
	}

	return nil
}
