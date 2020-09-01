package pipelines

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/Mindgamesnl/GoPack/pipelines/common_actions"
)

func Make113Pipeline() {
	pipeline := loader.CreatePipeline("to 1.13", "work/113/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaRevision(4))
	// remove comments
	pipeline.AddForFileType("json", common_actions.RemoveComment())

	pipeline.SaveUntouched()
	loader.AddPipeline(&pipeline)
}
