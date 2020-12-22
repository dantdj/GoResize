// Package resizing includes common methods for resizing images.
package resizing

import (
	"fmt"
	"strings"
	"time"

	"github.com/discordapp/lilliput"
)

// Default encoding options for different filetypes.
var EncodeOptions = map[string]map[int]int{
	".jpeg": {lilliput.JpegQuality: 85},
	".png":  {lilliput.PngCompression: 7},
	".webp": {lilliput.WebpQuality: 85},
}

// Resizes images provided in a byte array to the specified outputWidth and outputHeight.
func ResizeImage(inputBuf []byte, outputWidth int, outputHeight int) ([]byte, error) {
	// The decoder instantiation performs some basic checks,
	// such as magic bytes checking to match some known formats.
	decoder, err := lilliput.NewDecoder(inputBuf)
	if err != nil {
		fmt.Printf("Failed to decode image: %s\n", err)
		return nil, err
	}
	defer decoder.Close()

	header, err := decoder.Header()
	if err != nil {
		fmt.Printf("Failed to read image header: %s\n", err)
		return nil, err
	}

	if decoder.Duration() != 0 {
		fmt.Printf("duration: %.2f s\n", float64(decoder.Duration())/float64(time.Second))
	}

	ops := lilliput.NewImageOps(8192)
	defer ops.Close()

	var outputImg []byte

	// If shrinking the file, use a buffer the size of the original image.
	// If increasing the size, use a maximum of 50MB.
	// This is done to try and conserve memory allocations.
	if header.Width() > outputWidth || header.Height() > outputHeight {
		outputImg = make([]byte, 50*1024*1024)
	} else {
		outputImg = make([]byte, len(inputBuf))
	}

	if outputWidth == 0 {
		outputWidth = header.Width()
	}
	if outputHeight == 0 {
		outputHeight = header.Height()
	}

	resizeMethod := lilliput.ImageOpsResize

	if outputWidth == header.Width() && outputHeight == header.Height() {
		resizeMethod = lilliput.ImageOpsNoResize
	}

	outputType := "." + strings.ToLower(decoder.Description())

	opts := &lilliput.ImageOptions{
		FileType:             outputType,
		Width:                outputWidth,
		Height:               outputHeight,
		ResizeMethod:         resizeMethod,
		NormalizeOrientation: true,
		EncodeOptions:        EncodeOptions[outputType],
	}

	outputImg, err = ops.Transform(decoder, opts, outputImg)
	if err != nil {
		fmt.Printf("Failed to transform image, %s\n", err)
		return nil, err
	}

	return outputImg, nil
}
