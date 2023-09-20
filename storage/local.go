package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"go_grpc/lib/logger"
)

type Local struct {
	Directory string // relative public
}

func (m *Local) UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader) error {
	path := filepath.Join(m.Directory, bucketName, fileName)
	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logger.Error(ctx, "failed open local file", map[string]interface{}{
			"error": err,
			"tags":  []string{"storage", "local"},
		})

		return err
	}

	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		logger.Error(ctx, "failed copy local file", map[string]interface{}{
			"error": err,
			"tags":  []string{"storage", "local"},
		})

		return err
	}

	if fileInfo, err := os.Stat(path); fileInfo.Size() > 10000000 {
		if err != nil {
			return errors.New("file not found")
		}

		return errors.New("file too large")
	}

	return nil
}

func (m *Local) GetFileTemporaryURL(ctx context.Context, bucketName, filename string) (string, error) {
	return fmt.Sprintf("%s/%s", os.Getenv("CDN_BASE_URL"), filename), nil
}
