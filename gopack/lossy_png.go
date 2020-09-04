package gopack

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif" // for image.Decode() format registration
	_ "image/jpeg"
	"image/png"
	"os"
)

func CompressAsset(path string, strength int) []byte {
	return optimizePath(path, NoConversion, strength)
}

func optimizePath(inPath string, colorConversion ColorConversion, quantization int) []byte {
	// load image
	inFile, openErr := os.Open(inPath)
	if openErr != nil {
		fmt.Printf("couldn't open %v: %v\n", inPath, openErr)
		panic(openErr)
	}

	decoded, _, decodeErr := image.Decode(inFile)
	inFile.Close()
	if decodeErr != nil {
		fmt.Printf("couldn't decode %v: %v\n", inPath, decodeErr)
		panic(decodeErr)
	}

	optimized := Compress(decoded, colorConversion, quantization)

	content := new(bytes.Buffer)
	fuck := png.Encode(content, optimized)

	if fuck != nil {
		panic(fuck)
	}

	imageByets := content.Bytes()

	return imageByets
}