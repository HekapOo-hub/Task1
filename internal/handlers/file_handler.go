package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// FileHandler implements download and upload file for echo server
type FileHandler struct {
	fileService service.FileService
}

// Download is used for storing a file and showing it
// @Summary download file from filesystem
// @Security ApiKeyAuth
// @Tags file
// @Description to download file from filesystem
// @Param name path string true "filename"
// @Success 200 body png
// @Failure 400 body echo.NewHTTPError
// @Router /user/file/download/{name} [get]
func (f *FileHandler) Download(c echo.Context) error {
	fileName := c.Param("fileName")
	file, err := f.fileService.Download(c.Request().Context(), fileName)
	if err != nil {
		log.Warnf("file downloading error: %v", err)
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	dst, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		return fmt.Errorf("create file error in upload %w", err)
	}
	if _, err = io.Copy(dst, file); err != nil {
		return err
	}
	err = dst.Close()
	defer func() {
		if err := os.Remove(filepath.Clean(fileName)); err != nil {
			log.Warnf("remove file error in download %v", err)
		}
	}()

	return c.Attachment(dst.Name(), fileName)
}

// Upload is used for saving a file which was previously downloaded
// @Summary upload file
// @Security ApiKeyAuth
// @Tags file
// @Description to upload file from local variable which was previously downloaded
// @Param name path string true "filename"
// @Success 200 body string
// @Failure 400 body echo.NewHTTPError
// @Router /user/file/upload/{name} [get]
func (f *FileHandler) Upload(c echo.Context) error {
	fileName := c.Param("fileName")
	file, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		return fmt.Errorf("open file error in upload %w", err)
	}
	err = f.fileService.Upload(c.Request().Context(), file)
	if err != nil {
		log.Warnf("file uploading error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "file was uploaded!")
}
