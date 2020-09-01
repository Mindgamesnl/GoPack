package pipelines

import (
	"github.com/Mindgamesnl/GoPack/gopack"
	"github.com/Mindgamesnl/GoPack/gopack/common_actions"
)

func Make115Pipeline() {
	pipeline := gopack.CreatePipeline("to 1.15", "work/115/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaRevision(5))
	// remove comments
	pipeline.AddForFileType("json", common_actions.RemoveComment())

	pipeline.SaveUntouched()
	gopack.AddPipeline(&pipeline)
}
