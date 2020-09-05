package gopack

import "strings"

// don't compress highly visible assets
var compressionWhitelist = []string{
	"gui/",     // skip inventories etc
	"pack.png", // skip the pack icon
	".lang",    // skip language files
	"lang/",    // skip language files
}

func CompressResources(pipeline *Pipeline) {

	compressImage := func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// don't compress whitelist assets
		for i := range compressionWhitelist {
			if strings.Contains(resource.Path, compressionWhitelist[i]) {
				return
			}
		}
		fallbackData := resource.GetPipelineContent(pipeline)
		compressedData := CompressAsset(resource.OsPath, 20)
		if len(compressedData) <= 50 {
			pipeline.SaveBytes(resource, fallbackData)
		} else {
			pipeline.SaveBytes(resource, compressedData)
		}
	}

	compressJson := func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// don't compress whitelist assets
		for i := range compressionWhitelist {
			if strings.Contains(resource.Path, compressionWhitelist[i]) {
				return
			}
		}
		originalData := resource.GetPipelineString(pipeline)

		originalData = strings.Replace(originalData, "\r\n", "", -1)
		originalData = strings.Replace(originalData, "\n", "", -1)
		originalData = strings.Replace(originalData, " ", "", -1)
		originalData = strings.Replace(originalData, "	", "", -1)

		pipeline.SaveFile(resource, originalData)
	}

	pipeline.AddForFileType("json", compressJson)
	pipeline.AddForFileType("png", compressImage)
}
