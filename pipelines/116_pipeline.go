package pipelines

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/Mindgamesnl/GoPack/pipelines/common_actions"
)

func Make116Pipeline() {
	pipeline := loader.CreatePipeline("to 1.16", "work/116/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaRevision(6))

	pipeline.UnhandledFileHandler(func(originalPack loader.ResourcePack, resource loader.Resource, pipeline *loader.Pipeline) {
		// remove comments
		pipeline.SaveBytes(resource, resource.ContentAsBytes())
	})

	loader.AddPipeline(&pipeline)
}
