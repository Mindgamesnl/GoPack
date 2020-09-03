package gopack

import (
	"encoding/json"
	"strings"
)

func MigrateLanguage(pipeline *Pipeline, set map[string]string) {

	converter := func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// translate
		ogContent := resource.GetPipelineString(pipeline)
		delete(pipeline.WriteQueue, pipeline.OutFolder+resource.Path)

		elements := make(map[string]string)

		lines := strings.Split(ogContent, "\n")

		for i := range lines {
			line := lines[i]
			for s := range set {
				oldString := s
				updatedString := set[s]
				line = strings.Replace(line, oldString, updatedString, 1)
			}

			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				elements[parts[0]] = parts[1]
			}
		}

		resource.Path = strings.Replace(resource.Path, ".lang", ".json", 1)
		resource.ReadableName = strings.Replace(resource.ReadableName, ".lang", ".json", 1)
		resource.UniqueName = strings.Replace(resource.UniqueName, ".lang", ".json", 1)

		a, _ := json.Marshal(elements)

		pipeline.SaveBytes(resource, a)
	}

	// rename these files
	pipeline.AddForFileType("lang", converter)
}
