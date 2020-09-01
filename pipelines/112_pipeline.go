package pipelines

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/Mindgamesnl/GoPack/pipelines/common_actions"
)

func Make112Pipeline() {
	pipeline := loader.CreatePipeline("to 1.12.2", "work/112/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaValue(3))

	pipeline.UnhandledFileHandler(func(originalPack loader.ResourcePack, resource loader.Resource, pipeline *loader.Pipeline) {
		// remove comments
		pipeline.SaveBytes(resource, resource.ContentAsBytes())
	})

	loader.AddPipeline(&pipeline)
}
