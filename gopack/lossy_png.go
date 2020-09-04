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
	"path"
	"strings"
)

func CompressAsset(path string, strength int) []byte {
	return optimizePath(path, NoConversion, strength, "png", true)
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

func pathWithSuffix(filePath string, suffix string) string {
	extension := path.Ext(filePath)
	insertion := len(extension)
	if insertion > 0 {
		// if extension exists, trim it off of the base filename
		insertion = strings.LastIndex(filePath, extension)
	} else {
		insertion = len(filePath)
	}
	return filePath[:insertion] + suffix
}

func sizeDesc(size int64) string {
	suffixes := []string{"B", "kB", "MB", "GB", "TB"}
	var i int
	for i = 0; i+1 < len(suffixes); i++ {
		if size < 10000 {
			break
		}
		size = (size + 500) / 1000
	}
	return fmt.Sprintf("%d%v", size, suffixes[i])
}
