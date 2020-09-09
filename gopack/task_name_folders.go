package gopack

import "strings"

func RenamePackFolders(pipeline *Pipeline)  {
	// flattening part one, rename `textures/blocks` to `/textures/block`
	pipeline.AddGlobalHandler(func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		if strings.Contains(resource.Path, TransPath("textures/blocks")) {
			delete(pipeline.WriteQueue, pipeline.OutFolder+resource.Path)
			resource.Path = strings.Replace(resource.Path, TransPath("textures/blocks"), TransPath("textures/block"), 1)

			// overwrite handler, since we did handle it and don't want the original to be saved because that would be
			// silly, since we'd end up with both block and blocks
			pipeline.ProcessedFileNames = append(pipeline.ProcessedFileNames, resource.UniqueName)

			// save file manually with its new name
			pipeline.SaveBytes(resource, resource.GetPipelineContent(pipeline))
		}
	})

	// flattening part two, rename `textures/items` to `/textures/item`
	pipeline.AddGlobalHandler(func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		if strings.Contains(resource.Path, TransPath("textures/items")) {
			delete(pipeline.WriteQueue, pipeline.OutFolder+resource.Path)
			resource.Path = strings.Replace(resource.Path, TransPath("textures/items"), TransPath("textures/item"), 1)

			// overwrite handler, since we did handle it and don't want the original to be saved because that would be
			// silly, since we'd end up with both item and items
			pipeline.ProcessedFileNames = append(pipeline.ProcessedFileNames, resource.UniqueName)

			// save file manually with its new name
			pipeline.SaveBytes(resource, resource.GetPipelineContent(pipeline))
		}
	})
}
