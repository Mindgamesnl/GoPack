package gopack

import (
	"github.com/tidwall/sjson"
)

func RemoveComment() func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
	return func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		updatedJson, _ := sjson.Delete(resource.GetPipelineString(pipeline), "__comment")

		pipeline.SaveFile(resource, updatedJson)
	}
}
