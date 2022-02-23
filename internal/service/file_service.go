package service

import (
	"context"
	"fmt"
	"os"
)

// FileService is used for uploading and downloading files
type FileService struct {
	files []*os.File
}

// Download an image from filesystem
func (f *FileService) Download(ctx context.Context, fileName string) (*os.File, error) {
	for _, val := range f.files {
		if val.Name() == fileName {
			return val, nil
		}
	}
	return nil, fmt.Errorf("file not found")
}

// Upload an image to filesystem
func (f *FileService) Upload(ctx context.Context, file *os.File) error {
	if f.files == nil {
		f.files = make([]*os.File, 0)
	}
	f.files = append(f.files, file)
	return nil
}
