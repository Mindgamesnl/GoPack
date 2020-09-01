package common_actions

import (
	"github.com/Mindgamesnl/GoPack/gopack"
)

func SetMetaRevision(version int) func(originalPack gopack.ResourcePack, resource gopack.Resource, pipeline *gopack.Pipeline) {
	return func(originalPack gopack.ResourcePack, resource gopack.Resource, pipeline *gopack.Pipeline) {
		meta := &gopack.PackMcMeta{}
		gopack.JsonToStruct(resource.GetPipelineString(pipeline), meta)
		meta.Pack.PackFormat = version

		newContent := gopack.StructToJson(meta)

		pipeline.SaveFile(resource, newContent)
	}
}
