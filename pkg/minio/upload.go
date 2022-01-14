package minio

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"io"

	"github.com/minio/minio-go/v7"
	"log"
	"strings"
)

type Files struct {
	Context  context.Context
	Client   *minio.Client
	Bucket   string
	FileName string
	FileType string
	Reader   io.Reader
}

func NewFilesUploader(bucket string, fileName string, fileType string, reader io.Reader) *Files {
	client, err := InitMinIO()
	if err != nil {
		return nil
	}
	return &Files{Context: context.Background(), Client: client, Bucket: bucket, FileName: fileName, FileType: fileType, Reader: reader}
}

func (m *Files) Uploader() (*minio.UploadInfo, error) {
	reps := strings.Replace(m.FileName, " ", "-", -1)
	n := fmt.Sprintf("%s-%s", uuid.New().String(), reps)
	info, err := m.Client.PutObject(m.Context, m.Bucket, n, m.Reader, -1, minio.PutObjectOptions{ContentType: m.FileType})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("successfully uploaded %s of size %d\n", m.FileName, info.Size)
	return &info, nil
}

type File interface {
	Uploader() (*minio.UploadInfo, error)
}
