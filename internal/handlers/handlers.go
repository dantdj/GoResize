package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/dantdj/GoResize/pkg/resizing"
	"github.com/labstack/echo/v4"
)

// Default encoding options for different filetypes.
var contentTypes = map[string]string{
	".jpeg": "image/jpeg",
	".jpg":  "image/jpeg",
	".png":  "image/png",
	".webp": "image/webp",
}

func ResizeHandler(ctx echo.Context) error {
	formData, _ := ctx.FormFile("file")
	newWidth, _ := strconv.Atoi(ctx.QueryParam("width"))
	newHeight, _ := strconv.Atoi(ctx.QueryParam("height"))

	image, _ := formData.Open()
	defer image.Close()

	bytes, _ := ioutil.ReadAll(image)

	resized, err := resizing.ResizeImage(bytes, newWidth, newHeight)
	if err != nil {
		return ctx.NoContent(500)
	}

	return ctx.Blob(200, contentTypes[filepath.Ext(formData.Filename)], resized)
}

func HealthCheckHandler(ctx echo.Context) error {
	// A very simple health check.
	return ctx.String(http.StatusOK, `{"alive": true}`)
}
