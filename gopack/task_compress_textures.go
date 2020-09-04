package gopack

func CompressAssets(pipeline *Pipeline)  {

	pipeline.AddForFileType("png", func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		compressedData := CompressAsset(resource.OsPath, 20)

		// if its empty, try again with a lower strength
		if len(compressedData) == 0 {
			compressedData = CompressAsset(resource.OsPath, 10)
		}

		// 0 again? try 5
		if len(compressedData) == 0 {
			compressedData = CompressAsset(resource.OsPath, 5)
		}

		// 0 again? try 1
		if len(compressedData) == 0 {
			compressedData = CompressAsset(resource.OsPath, 1)
		}

		// still 0? log and don't do anything
		if len(compressedData) != 0 {
			pipeline.SaveBytes(resource, compressedData)
		} else {
			// pretend that we never did anything with it
			pipeline.ProcessedFileNames = RemoveIndex(pipeline.ProcessedFileNames, indexOf(resource.UniqueName, pipeline.ProcessedFileNames))
		}
	})

}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}