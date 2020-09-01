package gopack

func SetMetaRevision(version int) func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
	return func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		meta := &PackMcMeta{}
		JsonToStruct(resource.GetPipelineString(pipeline), meta)
		meta.Pack.PackFormat = version

		newContent := StructToJson(meta)

		pipeline.SaveFile(resource, newContent)
	}
}
