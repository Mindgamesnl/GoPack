package gopack

import "strings"

func ApplyFlatteningUpdate(pipeline *Pipeline) {

	// https://blockbench.net/2018/07/18/changes-to-resource-packs-in-minecraft-1-13/

	// flattening part one, rename `textures/blocks` to `/textures/block`
	pipeline.AddGlobalHandler(func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		if strings.Contains(resource.Path, "textures/blocks") {
			resource.Path = strings.Replace(resource.Path, "textures/blocks", "textures/block", 1)

			// overwrite handler, since we did handle it and don't want the original to be saved because that would be
			// silly, since we'd end up with both block and blocks
			pipeline.ProcessedFileNames = append(pipeline.ProcessedFileNames, resource.UniqueName)

			// save file manually with its new name
			pipeline.SaveBytes(resource, resource.GetPipelineContent(pipeline))
		}
	})

	// flattening part two, rename `textures/items` to `/textures/item`
	pipeline.AddGlobalHandler(func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		if strings.Contains(resource.Path, "textures/items") {
			resource.Path = strings.Replace(resource.Path, "textures/items", "textures/item", 1)

			// overwrite handler, since we did handle it and don't want the original to be saved because that would be
			// silly, since we'd end up with both item and items
			pipeline.ProcessedFileNames = append(pipeline.ProcessedFileNames, resource.UniqueName)

			// save file manually with its new name
			pipeline.SaveBytes(resource, resource.GetPipelineContent(pipeline))
		}
	})

	colors := [...]string{
		"white",
		"orange",
		"magenta",
		"light_blue",
		"yellow",
		"lime",
		"pink",
		"gray",
		"light_gray",
		"cyan",
		"purple",
		"blue",
		"brown",
		"green",
		"red",
		"black",
	}

	for i := range colors {
		color := colors[i]
		// wool
		pipeline.AddPathContainsHandler("wool_colored_" + color, rename("wool_colored_" + color, color + "_wool"))
		// stained glass and pane
		pipeline.AddPathContainsHandler("glass_" + color, rename("glass_" + color, color + "_stained_glass"))
		pipeline.AddPathContainsHandler("glass_pane_top_" + color, rename("glass_pane_top_" + color, color + "_stained_glass_pane_top"))
		// terrecotta or whatever
		pipeline.AddPathContainsHandler("hardened_clay_stained_" + color, rename("hardened_clay_stained_" + color, color + "_terracotta"))
		// concrete powder! how fun, fuck
		pipeline.AddPathContainsHandler("concrete_powder_" + color, rename("concrete_powder_" + color, color + "_concrete_powder"))
		// and regular concrete, aren't we lucky today
		pipeline.AddPathContainsHandler("concrete_" + color, rename("concrete_" + color, color + "_concrete"))
		// glazed terrecotta, the infinite joy
		pipeline.AddPathContainsHandler("glazed_terracotta_" + color, rename("glazed_terracotta_" + color, color + "_glazed_terracotta"))
		// shulkers aye, ffs
		pipeline.AddPathContainsHandler("shulker_top_" + color, rename("shulker_top_" + color, color + "_shulker_box_top"))
	}

	pipeline.AddPathContainsHandler("endercrystal", rename("endercrystal", "end_crystal"))

}

func rename(from string, to string) func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
	return func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// set and apply new name
		resource.Path = strings.Replace(resource.Path, from, to, 1)
		resource.ReadableName = strings.Replace(resource.ReadableName, from, to, 1)
		resource.UniqueName = strings.Replace(resource.UniqueName, from, to, 1)
		pipeline.SaveBytes(resource, resource.GetPipelineContent(pipeline))
	}
}
