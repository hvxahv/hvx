package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"log"
	"strings"
)

func (m *Minio) Uploader() (string, error) {
	err := m.NewBucket()
	if err != nil {
		log.Println("error in make bucket:", err)
		return "", err
	}

	file, err := m.File.Open()
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer file.Close()

	reps := strings.Replace(m.File.Filename, " ", "-", -1)
	fn := fmt.Sprintf("%s-%s", uuid.New().String(), reps)
	info, err := m.Client.PutObject(m.Ctx, m.Bucket, fn, file, -1, minio.PutObjectOptions{ContentType: m.FileType})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	log.Printf("successfully uploaded %s of size %d\n", m.File.Filename, info.Size)
	return fn, nil
}
