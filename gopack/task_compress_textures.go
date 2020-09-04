package gopack

func CompressAssets(pipeline *Pipeline)  {

	pipeline.AddForFileType("png", func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		fallbackData := resource.GetPipelineContent(pipeline)
		compressedData := CompressAsset(resource.OsPath, 30)

		// still 0? log and don't do anything
		if len(compressedData) <= 50 {
			pipeline.SaveBytes(resource, fallbackData)
		} else {
			pipeline.SaveBytes(resource, compressedData)
		}
	})

}

