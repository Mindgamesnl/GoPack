package common_actions

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/Mindgamesnl/GoPack/packs"
	"github.com/Mindgamesnl/GoPack/utils"
)

func SetMetaValue(version int) func(originalPack loader.ResourcePack, resource loader.Resource, pipeline *loader.Pipeline) {
	return func(originalPack loader.ResourcePack, resource loader.Resource, pipeline *loader.Pipeline) {
		meta := &packs.PackMcMeta{}
		utils.JsonToStruct(resource.ContentAsString(), meta)
		meta.Pack.PackFormat = version

		newContent := utils.StructToJson(meta)

		pipeline.SaveFile(resource, newContent)
	}
}
