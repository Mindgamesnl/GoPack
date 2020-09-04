package gopack

import "strings"

// don't compress highly visible assets
var compressionWhitelist = []string{
	"gui/",
	"pack.png",
}

func CompressResources(pipeline *Pipeline)  {

	compressImage := func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {

		// don't compress whitelist assets
		for i := range compressionWhitelist {
			if strings.Contains(resource.Path, compressionWhitelist[i]) {
				return
			}
		}

		fallbackData := resource.GetPipelineContent(pipeline)
		compressedData := CompressAsset(resource.OsPath, 30)

		// still 0? log and don't do anything
		if len(compressedData) <= 50 {
			pipeline.SaveBytes(resource, fallbackData)
		} else {
			pipeline.SaveBytes(resource, compressedData)
		}
	}

	pipeline.AddForFileType("png", compressImage)
	pipeline.AddForFileType("jpg", compressImage)
}
