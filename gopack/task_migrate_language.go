package gopack

import (
	"strings"
)

func MigrateLanguage(pipeline *Pipeline, set map[string]string) {

	converter := func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// translate
		ogContent := resource.GetPipelineString(pipeline)
		delete(pipeline.WriteQueue, pipeline.OutFolder+resource.Path)

		for s := range set {
			oldString := s
			updatedString := set[s]
			ogContent = strings.Replace(ogContent, oldString, updatedString, 1)
		}

		pipeline.SaveFile(resource, ogContent)
	}

	// rename these files
	pipeline.AddForFileType("lang", converter)
}
