package main

import (
	"github.com/Mindgamesnl/GoPack/gopack"
)

func main() {
	pack := gopack.FromZip("pack.zip")
	gopack.RegisterPipelines()
	gopack.RunPipelines(pack)
}
