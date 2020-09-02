package gopack

import (
	"strings"
)

func MakeEverythingLowercase(pipeline *Pipeline)  {
	// Force all file names to be LOWERCASE, regardless of their type or path
	// more info can be found in the last json migration of this update
	pipeline.AddGlobalHandler(func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// delete old file
		ogContent := resource.GetPipelineContent(pipeline)
		delete(pipeline.WriteQueue, pipeline.OutFolder+resource.Path)

		// replace and re-write
		resource.Path = strings.ToLower(resource.Path)
		resource.ReadableName = strings.ToLower(resource.ReadableName)
		resource.UniqueName = strings.ToLower(resource.UniqueName)
		pipeline.SaveBytes(resource, ogContent)
	})
}
