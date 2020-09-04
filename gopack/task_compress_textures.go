package gopack

func CompressAssets(pipeline *Pipeline)  {

	pipeline.AddForFileType("png", func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		compressedData := CompressAsset(resource.OsPath, 10)

		// 0 again? try 5
		if len(compressedData) == 0 {
			compressedData = CompressAsset(resource.OsPath, 5)
		}

		// 0 again? try 1
		if len(compressedData) == 0 {
			compressedData = CompressAsset(resource.OsPath, 1)
		}

		// 0 again? try 0, so just do colour
		if len(compressedData) == 0 {
			compressedData = CompressAsset(resource.OsPath, 0)
		}

		// still 0? log and don't do anything
		if len(compressedData) != 0 {
			pipeline.SaveBytes(resource, compressedData)
		}
	})

}
