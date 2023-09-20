package storage

import (
	"context"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"go_grpc/lib"
	"go_grpc/lib/logger"
)

type Minio struct {
	Client *lib.MinioClient
}

func (m *Minio) UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader) error {
	_, err := m.Client.PutObject(ctx, bucketName, fileName, file, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logger.Error(ctx, "failed upload file to minio", map[string]interface{}{
			"error": err,
			"tags":  []string{"storage", "minio"},
		})
	}

	return err
}

func (m *Minio) GetFileTemporaryURL(ctx context.Context, bucketName, filename string) (string, error) {
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)

	// Generates a presigned url which expires in a day.
	presignedURL, err := m.Client.PresignedGetObject(ctx, bucketName, filename, time.Second*24*60*60, reqParams)
	if err != nil {
		logger.Error(ctx, "failed get file temporary url", map[string]interface{}{
			"error": err,
			"tags":  []string{"storage", "minio"},
		})

		return "", err
	}

	return os.Getenv("CDN_BASE_URL") + presignedURL.Path + "?" + presignedURL.RawQuery, nil
}
