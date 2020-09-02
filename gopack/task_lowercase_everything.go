package gopack

import "strings"

func MakeEverythingLowercase(pipeline *Pipeline)  {
	// Force all file names to be LOWERCASE, regardless of their type or path
	// more info can be found in the last json migration of this update
	pipeline.AddGlobalHandler(func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		resource.Path = strings.ToLower(resource.Path)
		resource.ReadableName = strings.ToLower(resource.ReadableName)
		resource.UniqueName = strings.ToLower(resource.UniqueName)
		pipeline.SaveBytes(resource, resource.GetPipelineContent(pipeline))
	})
}
