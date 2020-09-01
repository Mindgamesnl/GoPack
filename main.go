package main

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/sirupsen/logrus"
)

func main() {
	originalPack := loader.FromZip("pack.zip")
	logrus.Info(originalPack)


}
