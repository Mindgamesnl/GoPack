package gopack

import (
	"strings"
)

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

	// renaming a fuck ton of blocks, because mojang just likes to do that sometimes
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

	// more renaming, because mojang just loves that
	pipeline.AddPathContainsHandler("endercrystal", rename("endercrystal", "end_crystal"))

	// EVEN MORE!
	// MINECRAFT IS THE FUCKING GIFT THAT JUST KEEPS ON GIVING FOLKS
	// BECAUSE WE GOT A MILLION AND ONE TYPES OF STONE THAT NEED CONVERTION TOO!
	// WHOHOOO!
	// my life is a meme
	pipeline.AddPathContainsHandler("stone_andesite", rename("stone_andesite", "andesite"))
	pipeline.AddPathContainsHandler("stone_andesite_smooth", rename("stone_andesite_smooth", "polished_andesite"))
	pipeline.AddPathContainsHandler("stone_diorite", rename("stone_diorite", "diorite"))
	pipeline.AddPathContainsHandler("stone_diorite_smooth", rename("stone_diorite_smooth", "polished_diorite"))
	pipeline.AddPathContainsHandler("stone_granite", rename("stone_granite", "granite"))
	pipeline.AddPathContainsHandler("stone_granite_smooth", rename("stone_granite_smooth", "polished_granite"))

	// even more stone..
	pipeline.AddPathContainsHandler("cobblestone_mossy", rename("cobblestone_mossy", "mossy_cobblestone"))
	pipeline.AddPathContainsHandler("stonebrick", rename("stonebrick", "stone_bricks"))
	pipeline.AddPathContainsHandler("stonebrick_carved", rename("stonebrick_carved", "chiseled_stone_bricks"))
	pipeline.AddPathContainsHandler("stonebrick_cracked", rename("stonebrick_cracked", "cracked_stone_bricks"))
	pipeline.AddPathContainsHandler("stonebrick_mossy", rename("stonebrick_mossy", "mossy_stone_bricks"))
	pipeline.AddPathContainsHandler("stonebrick_mossy", rename("stonebrick_mossy", "mossy_stone_bricks"))

	// sand stone, ofcource, yes sure, fuck you
	pipeline.AddPathContainsHandler("sandstone_normal", rename("sandstone_normal", "sandstone"))
	pipeline.AddPathContainsHandler("sandstone_carved", rename("sandstone_carved", "chiseled_sandstone"))
	pipeline.AddPathContainsHandler("sandstone_smooth", rename("sandstone_smooth", "cut_sandstone"))
	pipeline.AddPathContainsHandler("sandstone_smooth", rename("sandstone_smooth", "cut_sandstone"))

	// lets do that again but for red sand stone, the colour of fucking blood, ew epic gamer moment right there
	pipeline.AddPathContainsHandler("red_sandstone_normal", rename("red_sandstone_normal", "red_sandstone"))
	pipeline.AddPathContainsHandler("red_sandstone_carved", rename("red_sandstone_carved", "chiseled_red_sandstone"))
	pipeline.AddPathContainsHandler("red_sandstone_smooth", rename("red_sandstone_smooth", "cut_red_sandstone"))
	pipeline.AddPathContainsHandler("red_sandstone_smooth", rename("red_sandstone_smooth", "cut_red_sandstone"))

	// dirt! as a epic in game reference to the human dirt that is Notch, because i'm getting sick of this format

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
