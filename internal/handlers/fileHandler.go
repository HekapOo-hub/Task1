package handlers

import (
	"net/http"

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
// @Router /file/download/{name} [get]
func (f *FileHandler) Download(c echo.Context) error {
	fileName := c.Param("fileName")
	err := f.fileService.Download(c.Request().Context(), fileName)
	if err != nil {
		log.Warnf("file downloading error: %v", err)
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return c.File(fileName)
}

// Upload is used for saving a file which was previously downloaded
// @Summary upload file
// @Security ApiKeyAuth
// @Tags file
// @Description to upload file from local variable which was previously downloaded
// @Param name path string true "filename"
// @Success 200 body string
// @Failure 400 body echo.NewHTTPError
// @Router /file/upload/{name} [get]
func (f *FileHandler) Upload(c echo.Context) error {
	fileName := c.Param("fileName")
	err := f.fileService.Upload(c.Request().Context(), fileName)
	if err != nil {
		log.Warnf("file uploading error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "file was uploaded!")
}
