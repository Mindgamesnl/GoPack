package gopack

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func readableSize(size int) string {
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

func WriteToFile(targetFolder string, content []byte)  {
	fa := os.MkdirAll(filepath.Dir(targetFolder), os.ModePerm)
	if fa != nil {
		panic(fa)
	}

	f, err := os.Create(targetFolder)

	if err != nil {
		panic(err)
	}

	defer f.Close()
	w := bufio.NewWriter(f)
	_, err2 := w.Write(content)

	if err2 != nil {
		panic(err2)
	}
	fuck := w.Flush()
	if fuck != nil {
		panic(fuck)
	}
}