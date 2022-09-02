package fs

import (
	"fmt"
	"github.com/hvxahv/hvx/errors"
	"io"
)

type Bucket interface {
	// Create a storage bucket and use the provider to distinguish the fs server.
	Create() error
}

type Files interface {
	// Put method and use the provider to distinguish the fs server.
	Put() (string, error)

	// Delete objects from the object store by filename.
	Delete() error
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
		var policy string
		switch f.BucketName {
		case "avatar":
			policy = MinioPublicAvatarBucketPolicy
		case "attach":
			policy = MinioPublicAttachBucketPolicy
		default:
			return errors.New("don't know the name to create the bucket.")
		}
		if err := client.MakeBucket(f.BucketName, policy); err != nil {
			return err
		}
	default:
		return errors.New(errors.ErrFSProvider)
	}
	return nil
}

func NewFsPut(provider string, bucketName string, fileName string, contentType string, buffer io.Reader, fileSize int64) *Fs {
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

func NewFsDelete(provider string, bucketName string, fileName string) *Fs {
	return &Fs{Provider: provider, BucketName: bucketName, FileName: fileName}
}

func (f *Fs) Delete() error {
	switch f.Provider {
	case "minio":
		minio, err := NewDefaultMinio().Dial()
		if err != nil {
			return err
		}
		if err := NewMinioRemoveFile(minio.Client, minio.Ctx, f.BucketName, f.FileName).Remove(); err != nil {
			return err
		}
		return nil

	default:
		return errors.New("NO_MATCHING_PROVIDER")
	}
}
