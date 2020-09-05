package gopack

import (
	"strings"
)

func ForceContent(pipeline *Pipeline, set map[string]string) {
	for s := range set {
		oldName := s
		content := set[s]

		converter := func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
			// don't save original
			if strings.HasSuffix(resource.Path, oldName) {
				delete(pipeline.WriteQueue, pipeline.OutFolder+resource.Path)
				pipeline.SaveFile(resource, content)
			}
		}

		// rename these files
		pipeline.AddGlobalHandler(converter)
	}
}
