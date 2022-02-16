package service

import (
	"context"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func createImage(fileName string) (*os.File, error) {
	width := 200
	height := 100

	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{R: 100, G: 200, B: 200, A: 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	err = png.Encode(f, img)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func TestFileService_Download(t *testing.T) {
	fileName := "my.png"
	_, err := createImage(fileName)
	require.NoError(t, err)
	fileService := FileService{file: &os.File{}}
	err = fileService.Download(context.Background(), fileName)
	require.NotEmpty(t, fileService.file)
	require.NoError(t, err)
	err = os.Remove(filepath.Clean(fileName))
	require.NoError(t, err)
}

func TestFileService_Upload(t *testing.T) {
	fileName := "cpy.png"
	_, err := createImage(fileName)
	require.NoError(t, err)
	file, err := os.Open(fileName)
	require.NoError(t, err)
	fileService := FileService{file: file}
	err = fileService.Upload(context.Background(), "upload.png")
	require.NoError(t, err)
	err = os.Remove(filepath.Clean(fileName))
	require.NoError(t, err)
}
