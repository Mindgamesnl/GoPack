package common_actions

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/tidwall/sjson"
)

func RemoveComment() func(originalPack loader.ResourcePack, resource loader.Resource, pipeline *loader.Pipeline) {
	return func(originalPack loader.ResourcePack, resource loader.Resource, pipeline *loader.Pipeline) {
		updatedJson, _ := sjson.Delete(resource.GetPipelineString(pipeline), "__comment")

		pipeline.SaveFile(resource, updatedJson)
	}
}
