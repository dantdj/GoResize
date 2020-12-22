package handlers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dantdj/GoResize/pkg/resizing"
	"github.com/labstack/echo/v4"
)

func ResizeHandler(ctx echo.Context) error {
	formData, _ := ctx.FormFile("file")
	newWidth, _ := strconv.Atoi(ctx.QueryParam("width"))
	newHeight, _ := strconv.Atoi(ctx.QueryParam("height"))

	image, _ := formData.Open()
	defer image.Close()

	bytes, _ := ioutil.ReadAll(image)

	resized, _ := resizing.ResizeImage(bytes, newWidth, newHeight)

	return ctx.Blob(200, "image/png", resized)
}

func HealthCheckHandler(ctx echo.Context) error {
	// A very simple health check.
	return ctx.String(http.StatusOK, `{"alive": true}`)
}
