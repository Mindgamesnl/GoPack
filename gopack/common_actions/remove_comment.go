package common_actions

import (
	"github.com/Mindgamesnl/GoPack/gopack"
	"github.com/tidwall/sjson"
)

func RemoveComment() func(originalPack gopack.ResourcePack, resource gopack.Resource, pipeline *gopack.Pipeline) {
	return func(originalPack gopack.ResourcePack, resource gopack.Resource, pipeline *gopack.Pipeline) {
		updatedJson, _ := sjson.Delete(resource.GetPipelineString(pipeline), "__comment")

		pipeline.SaveFile(resource, updatedJson)
	}
}
