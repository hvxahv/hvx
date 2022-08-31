package fs

import (
	"fmt"
	"github.com/hvxahv/hvx/errors"
	"io"
)

type Bucket interface {
	Create() error
}

type Files interface {
	Put() (string, error)
}

type Fs struct {
	Provider   string
	BucketName string

	FileName    string
	ContentType string
	FileBuffer  io.Reader
	FileSize    int64
}

func NewBucket(provider string, bucketName string) *Fs {
	return &Fs{Provider: provider, BucketName: bucketName}
}

func (f *Fs) Create() error {
	switch f.Provider {
	case "minio":
		client, err := NewDefaultMinio().Dial()
		if err != nil {
			return err
		}
		if err := client.MakeBucket(f.BucketName, MinioPublicBucketPolicy); err != nil {
			return err
		}
	default:
		return errors.New("NO_MATCHING_PROVIDER")
	}
	return nil
}

func NewFs(provider string, bucketName string, fileName string, contentType string, buffer io.Reader, fileSize int64) *Fs {
	return &Fs{Provider: provider, BucketName: bucketName, FileName: fileName, ContentType: contentType, FileBuffer: buffer, FileSize: fileSize}
}

func (f *Fs) Put() (string, error) {
	switch f.Provider {
	case "minio":
		client, err := NewDefaultMinio().Dial()
		if err != nil {
			return "", err
		}
		put, err := NewMinioFiles(client.Client, client.Ctx, f.BucketName, f.FileName, f.ContentType, f.FileBuffer, f.FileSize).Put()
		if err != nil {
			return "", err
		}
		if client.useSSL {
			return fmt.Sprintf("%s://%s/%s/%s", "https", client.endpoint, put.Bucket, f.FileName), nil
		}
		return fmt.Sprintf("%s://%s/%s/%s", "http", client.endpoint, put.Bucket, f.FileName), nil

	default:
		return "", errors.New("NO_MATCHING_PROVIDER")
	}
}
