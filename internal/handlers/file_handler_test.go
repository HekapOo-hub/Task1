package handlers

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestFileHandler_Upload(t *testing.T) {
	file, err := createImage("upload.png")
	request, err := http.NewRequest(http.MethodGet,
		url+"user/file/upload/"+file.Name(), nil)
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test file handler upload error %v", err)
		}
	}()
	responseBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "file was uploaded!", string(responseBody))
	err = os.Remove(filepath.Clean(file.Name()))
	require.NoError(t, err)
}

func TestFileHandler_Download(t *testing.T) {
	fileName := "myFile.png"
	request, err := http.NewRequest(http.MethodGet,
		url+"user/file/download/"+fileName, nil)
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	_, err = (&http.Client{}).Do(request)
	require.NoError(t, err)
}

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
