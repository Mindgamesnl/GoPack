package pipelines

import (
	"github.com/Mindgamesnl/GoPack/gopack"
	"github.com/Mindgamesnl/GoPack/gopack/common_actions"
)

func Make111Pipeline() {
	pipeline := gopack.CreatePipeline("to 1.11", "work/111/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaRevision(3))
	// remove comments
	pipeline.AddForFileType("json", common_actions.RemoveComment())

	pipeline.SaveUntouched()
	gopack.AddPipeline(&pipeline)
}
