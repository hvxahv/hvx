package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func (m *Minio) NewBucket() error {
	location := viper.GetString("minio.location")
	cli := m.Client

	if err := cli.MakeBucket(m.Ctx, m.Bucket, minio.MakeBucketOptions{Region: location}); err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := cli.BucketExists(m.Ctx, m.Bucket)
		if errBucketExists == nil && exists {
			// If the Bucket has been created, the output has already created the Bucket.
			return nil
		} else {
			return err
		}
	} else {
		// Set bucket permissions.
		// Set the access permissions of the bucket.
		// By default, it has read and write permissions so that the returned url can be
		// accessed and downloaded in the browser. It can be changed according to business needs.
		policy := `{"Version":"2012-10-17",
"Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},
"Action":["s3:GetBucketLocation","s3:ListBucket","s3:ListBucketMultipartUploads"],
"Resource":["arn:aws:s3:::avatar"]},{"Effect":"Allow","Principal":{"AWS":["*"]},
"Action":["s3:AbortMultipartUpload","s3:DeleteObject","s3:GetObject","s3:ListMultipartUploadParts","s3:PutObject"],
"Resource":["arn:aws:s3:::avatar/*"]}]}
`
		if err2 := cli.SetBucketPolicy(context.Background(), m.Bucket, policy); err2 != nil {
			return err2
		}

		// The method used to obtain minio permissions,
		// returns the detailed permissions of the bucket,
		// which can be used to obtain settings,
		// in development and used in the method ( SetBucketPolicy() ) of setting permissions.
		// p, err := client.GetBucketPolicy(context.Background(), m.bucket)
	}
	return nil
}
