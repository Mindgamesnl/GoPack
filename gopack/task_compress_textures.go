package gopack

func CompressAssets(pipeline *Pipeline)  {

	pipeline.AddForFileType("png", func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		compressedData := CompressAsset(resource.OsPath, 20)
		pipeline.SaveBytes(resource, compressedData)
	})

}
