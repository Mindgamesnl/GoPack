package pipelines

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/Mindgamesnl/GoPack/pipelines/common_actions"
)

func Make115Pipeline() {
	pipeline := loader.CreatePipeline("to 1.15", "work/115/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaRevision(5))

	pipeline.SaveUntouched()
	loader.AddPipeline(&pipeline)
}
