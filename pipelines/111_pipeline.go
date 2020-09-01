package pipelines

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/Mindgamesnl/GoPack/pipelines/common_actions"
)

func Make111Pipeline() {
	pipeline := loader.CreatePipeline("to 1.11", "work/111/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaRevision(3))
	// remove comments
	pipeline.AddForFileType("json", common_actions.RemoveComment())

	pipeline.SaveUntouched()
	loader.AddPipeline(&pipeline)
}
