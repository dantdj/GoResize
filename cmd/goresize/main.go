package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/discordapp/lilliput"
)

var EncodeOptions = map[string]map[int]int{
	".jpeg": map[int]int{lilliput.JpegQuality: 85},
	".png":  map[int]int{lilliput.PngCompression: 7},
	".webp": map[int]int{lilliput.WebpQuality: 85},
}

func main() {
	var inputFilename string
	var outputWidth int
	var outputHeight int

	flag.StringVar(&inputFilename, "input", "", "name of input file to resize/transcode")
	flag.IntVar(&outputWidth, "width", 0, "width of output file")
	flag.IntVar(&outputHeight, "height", 0, "height of output file")
	flag.Parse()

	inputBuf, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", err)
		os.Exit(1)
	}

	// The decoder instantiation performs some basic checks,
	// such as magic bytes checking to match some known formats.
	decoder, err := lilliput.NewDecoder(inputBuf)
	if err != nil {
		fmt.Printf("Failed to decode imge: %s\n", err)
		os.Exit(1)
	}
	defer decoder.Close()

	header, err := decoder.Header()
	if err != nil {
		fmt.Printf("Failed to read image header: %s\n", err)
		os.Exit(1)
	}

	if decoder.Duration() != 0 {
		fmt.Printf("duration: %.2f s\n", float64(decoder.Duration())/float64(time.Second))
	}

	// get ready to resize image,
	// using 8192x8192 maximum resize buffer size
	ops := lilliput.NewImageOps(8192)
	defer ops.Close()

	// Create a buffer to store the output image.
	// If shrinking the file, use a buffer the size of the original image.
	// If increasing the size, use a maximum of 50MB.
	// This is done to try and conserve memory allocations.
	var outputImg []byte

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
		os.Exit(1)
	}

	outputFilename := "resized" + filepath.Ext(inputFilename)

	if _, err := os.Stat(outputFilename); !os.IsNotExist(err) {
		fmt.Printf("Output filename %s exists, quitting\n", outputFilename)
		os.Exit(1)
	}

	err = ioutil.WriteFile(outputFilename, outputImg, 0400)
	if err != nil {
		fmt.Printf("Failed to write out resized image, %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Image written to %s\n", outputFilename)
}