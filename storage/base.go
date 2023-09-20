package storage

import (
	"context"
	"io"
)

type Storage interface {
	UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader) error
	GetFileTemporaryURL(ctx context.Context, bucketName, filename string) (string, error)
}
