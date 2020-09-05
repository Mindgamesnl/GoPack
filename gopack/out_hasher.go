package gopack

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

var hashes = make(map[string]string)

func HashFile(path string) {
	elements := strings.Split(path, "/")
	name := elements[len(elements) - 1]
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	hash := sha1.New()

	if _, err := io.Copy(hash, file); err != nil {
		panic(err)
	}

	// 20 byte hash
	hashInBytes := hash.Sum(nil)[:20]
	strHash := hex.EncodeToString(hashInBytes)
	hashes[name] = strHash
}

func SaveHashes()  {
	hashes := StructToJson(hashes)
	WriteToFile("out/hashes.json", []byte(hashes))
}