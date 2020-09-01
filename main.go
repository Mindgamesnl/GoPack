package main

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/Mindgamesnl/GoPack/pipelines"
)

func main() {
	pack := loader.FromZip("pack.zip")
	pipelines.RegisterPipelines()
	loader.RunPipelines(pack)
}
