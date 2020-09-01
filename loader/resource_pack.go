package loader

import (
	"github.com/Mindgamesnl/GoPack/utils"
	"github.com/sirupsen/logrus"
)

type ResourcePack struct {
	VersionString string
	VersionNumber int8
	Root string
	FileHandler FileCollection
}

func FromZip(filename string)  {
	files, _ := utils.Unzip(filename, "work/original/")
	for i := range files {
		logrus.Info(files[i])
	}
}