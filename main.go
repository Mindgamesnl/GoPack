package main

import (
	"github.com/Mindgamesnl/GoPack/gopack"
)

func main() {
	// hypixel.GenGoCode()
	pack := gopack.FromZip("pack.zip")
	gopack.RegisterPipelines()
	gopack.RunPipelines(pack)
}
