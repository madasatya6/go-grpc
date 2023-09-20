package lib

import (
	"github.com/minio/minio-go/v7"
)

type MinioClient struct {
	*minio.Client
}
