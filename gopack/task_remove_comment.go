package gopack

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"strings"
)

func RemoveComment() func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
	return func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		scan := FindJsonKeys(gjson.Parse(resource.GetPipelineString(pipeline)), resource.OsPath)

		updatedJson := resource.GetPipelineString(pipeline)

		for i := range scan {
			key := scan[i]
			if strings.Contains(key, "__comment") {
				var err error
				updatedJson, err = sjson.Delete(updatedJson, key)
				if err != nil {
					panic(err)
				}
			}
		}

		pipeline.SaveFile(resource, updatedJson)
	}
}
