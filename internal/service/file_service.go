package service

import (
	"context"
	"fmt"
	"os"
)

// FileService is used for uploading and downloading files
type FileService struct {
	file *os.File
}

// Download an image from filesystem
func (f *FileService) Download(ctx context.Context) (*os.File, error) {
	if f.file == nil {
		return nil, fmt.Errorf("file not found")
	}
	return f.file, nil
}

// Upload an image to filesystem
func (f *FileService) Upload(ctx context.Context, file *os.File) error {
	f.file = file
	return nil
}
