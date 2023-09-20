package config

import (
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"go_grpc/lib"
	"go_grpc/storage"
)

func NewMinioClient() *lib.MinioClient {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	username := os.Getenv("MINIO_USERNAME")
	password := os.Getenv("MINIO_PASSWORD")
	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_SSL"))
	if err != nil {
		panic(err)
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(username, password, ""),
		Secure: useSSL,
	})

	if err != nil {
		panic(err)
	}

	return &lib.MinioClient{Client: client}
}

func NewLocalStorage() *storage.Local {
	return &storage.Local{Directory: "public"}
}

func NewStorage() storage.Storage {
	if os.Getenv("FORCE_LOCAL_STORAGE") != "true" {
		minioClient := NewMinioClient()
		return &storage.Minio{Client: minioClient}
	}

	return NewLocalStorage()
}
