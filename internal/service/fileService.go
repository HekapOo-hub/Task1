package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FileService is used for uploading and downloading files
type FileService struct {
	file *os.File
}

// Download an image from filesystem
func (f *FileService) Download(ctx context.Context, fileName string) error {
	file, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		return fmt.Errorf("open file error in download %w", err)
	}
	f.file = file
	return nil
}

// Upload an image to filesystem
func (f *FileService) Upload(ctx context.Context, fileName string) error {
	dst, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		return fmt.Errorf("create file error in upload %w", err)
	}
	if _, err = io.Copy(dst, f.file); err != nil {
		return err
	}
	err = dst.Close()
	if err != nil {
		return fmt.Errorf("error in close "+fileName+" %w", err)
	}
	return nil
}
