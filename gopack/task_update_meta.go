package gopack

import "github.com/Mindgamesnl/GoPack/gopack/utils"

func SetMetaRevision(version int) func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
	return func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		meta := &PackMcMeta{}
		utils.JsonToStruct(resource.GetPipelineString(pipeline), meta)
		meta.Pack.PackFormat = version

		newContent := utils.StructToJson(meta)

		pipeline.SaveFile(resource, newContent)
	}
}
