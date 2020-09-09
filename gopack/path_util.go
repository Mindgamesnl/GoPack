package gopack

import (
	"os"
	"strings"
)

func TransPath(path string) string {
	o := strings.Replace(path,"/", string(os.PathSeparator), -1)
	return o
}