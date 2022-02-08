package handlers

import (
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// FileHandler implements download and upload file for echo server
type FileHandler struct {
	fileService service.FileService
}

// Download is used for storing a file and showing it
func (f *FileHandler) Download(c echo.Context) error {
	fileName := c.Param("fileName")
	err := f.fileService.Download(c.Request().Context(), fileName)
	if err != nil {
		log.WithField("error", err).Warn("file downloading error")
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return c.File(fileName)
}

// Upload is used for saving a file which was previously downloaded
func (f *FileHandler) Upload(c echo.Context) error {
	fileName := c.Param("fileName")
	err := f.fileService.Upload(c.Request().Context(), fileName)
	if err != nil {
		log.WithField("error", err).Warn("file uploading error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "file was uploaded!")
}
