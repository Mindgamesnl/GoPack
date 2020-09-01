package pipelines

import (
	"github.com/Mindgamesnl/GoPack/gopack"
	"github.com/Mindgamesnl/GoPack/gopack/common_actions"
)

func Make113Pipeline() {
	pipeline := gopack.CreatePipeline("to 1.13", "work/113/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaRevision(4))
	// remove comments
	pipeline.AddForFileType("json", common_actions.RemoveComment())

	pipeline.SaveUntouched()
	gopack.AddPipeline(&pipeline)
}
