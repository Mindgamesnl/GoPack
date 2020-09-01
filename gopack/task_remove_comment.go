package gopack

import (
	"github.com/Mindgamesnl/GoPack/gopack/utils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"strings"
)

func RemoveComment() func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
	return func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		scan := utils.FindJsonKeys(gjson.Parse(resource.GetPipelineString(pipeline)), "")

		updatedJson := resource.GetPipelineString(pipeline)

		for i := range scan {
			key := scan[i]
			if strings.Contains(key, "__comment") {
				updatedJson, _ = sjson.Delete(resource.GetPipelineString(pipeline), key)
			}
		}

		pipeline.SaveFile(resource, updatedJson)
	}
}
