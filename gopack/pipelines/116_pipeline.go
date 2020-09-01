package pipelines

import (
	"github.com/Mindgamesnl/GoPack/gopack"
	"github.com/Mindgamesnl/GoPack/gopack/common_actions"
)

func Make116Pipeline() {
	pipeline := gopack.CreatePipeline("to 1.16", "work/116/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaRevision(6))
	// remove comments
	pipeline.AddForFileType("json", common_actions.RemoveComment())

	pipeline.SaveUntouched()
	gopack.AddPipeline(&pipeline)
}
