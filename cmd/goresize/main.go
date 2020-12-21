package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dantdj/GoResize/pkg/resizing"
)

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

	outputImg, err := resizing.ResizeImage(inputBuf, outputWidth, outputHeight)
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
