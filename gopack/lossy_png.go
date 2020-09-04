package gopack

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	_ "image/gif" // for image.Decode() format registration
	_ "image/jpeg"
	"image/png"
	"os"
)

func CompressAsset(path string, strength int) []byte {
	return optimizePath(path, RGBAConversion, strength, "png", true)
}

func optimizePath(inPath string, colorConversion ColorConversion, quantization int, extension string, rewriteOriginal bool, ) []byte {
	// load image
	inFile, openErr := os.Open(inPath)
	if openErr != nil {
		fmt.Printf("couldn't open %v: %v\n", inPath, openErr)
		return nil
	}

	decoded, _, decodeErr := image.Decode(inFile)
	inFile.Close()
	if decodeErr != nil {
		fmt.Printf("couldn't decode %v: %v\n", inPath, decodeErr)
		return nil
	}

	optimized := Compress(decoded, colorConversion, quantization)

	var content bytes.Buffer
	bufferWriter := bufio.NewWriter(&content)
	fuck := png.Encode(bufferWriter, optimized)

	if fuck != nil {
		panic(fuck)
	}

	imageByets := content.Bytes()

	return imageByets
}