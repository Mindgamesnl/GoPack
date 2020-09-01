package pipelines

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/Mindgamesnl/GoPack/pipelines/common_actions"
)

func Make116Pipeline() {
	pipeline := loader.CreatePipeline("to 1.16", "work/116/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaRevision(6))
	// remove comments
	pipeline.AddForFileType("json", common_actions.RemoveComment())

	pipeline.SaveUntouched()
	loader.AddPipeline(&pipeline)
}
