package gopack

import (
	"strconv"
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
	pipeline.AddPathContainsHandler("stone_andesite_smooth", rename("stone_andesite_smooth", "polished_andesite"))
	pipeline.AddPathContainsHandler("stone_andesite", rename("stone_andesite", "andesite"))
	pipeline.AddPathContainsHandler("stone_diorite_smooth", rename("stone_diorite_smooth", "polished_diorite"))
	pipeline.AddPathContainsHandler("stone_diorite", rename("stone_diorite", "diorite"))
	pipeline.AddPathContainsHandler("stone_granite_smooth", rename("stone_granite_smooth", "polished_granite"))
	pipeline.AddPathContainsHandler("stone_granite", rename("stone_granite", "granite"))

	// even more stone..
	pipeline.AddPathContainsHandler("cobblestone_mossy", rename("cobblestone_mossy", "mossy_cobblestone"))
	pipeline.AddPathContainsHandler("stonebrick_carved", rename("stonebrick_carved", "chiseled_stone_bricks"))
	pipeline.AddPathContainsHandler("stonebrick_cracked", rename("stonebrick_cracked", "cracked_stone_bricks"))
	pipeline.AddPathContainsHandler("stonebrick_mossy", rename("stonebrick_mossy", "mossy_stone_bricks"))
	pipeline.AddPathContainsHandler("stonebrick_mossy", rename("stonebrick_mossy", "mossy_stone_bricks"))
	pipeline.AddPathContainsHandler("stonebrick", rename("stonebrick", "stone_bricks"))

	// sand stone, ofcource, yes sure, fuck you
	pipeline.AddPathContainsHandler("sandstone_normal", rename("sandstone_normal", "sandstone"))
	pipeline.AddPathContainsHandler("sandstone_carved", rename("sandstone_carved", "chiseled_sandstone"))
	pipeline.AddPathContainsHandler("sandstone_smooth", rename("sandstone_smooth", "cut_sandstone"))

	// lets do that again but for red sand stone, the colour of fucking blood, ew epic gamer moment right there
	pipeline.AddPathContainsHandler("red_sandstone_normal", rename("red_sandstone_normal", "red_sandstone"))
	pipeline.AddPathContainsHandler("red_sandstone_carved", rename("red_sandstone_carved", "chiseled_red_sandstone"))
	pipeline.AddPathContainsHandler("red_sandstone_smooth", rename("red_sandstone_smooth", "cut_red_sandstone"))

	// dirt! as a epic in game reference to the human dirt that is Notch, because i'm getting sick of this format
	pipeline.AddPathContainsHandler("grass_side_overlay", rename("grass_side_overlay", "grass_block_side_overlay"))
	pipeline.AddPathContainsHandler("grass_side_snowed", rename("grass_side_snowed", "grass_block_snow"))
	pipeline.AddPathContainsHandler("grass_top", rename("grass_top", "grass_block_top"))
	pipeline.AddPathContainsHandler("dirt_podzol_side", rename("dirt_podzol_side", "podzol_side"))
	pipeline.AddPathContainsHandler("dirt_podzol_top", rename("dirt_podzol_top", "podzol_top"))
	pipeline.AddPathContainsHandler("farmland_dry", rename("farmland_dry", "farmland"))
	pipeline.AddPathContainsHandler("grass_side", rename("grass_side", "grass_block_side"))

	// quartz! no hate here, i actually like quartz, used it in my first minecraft builds way back when
	pipeline.AddPathContainsHandler("quartz_block_chiseled_top", rename("quartz_block_chiseled_top", "chiseled_quartz_block_top"))
	pipeline.AddPathContainsHandler("quartz_block_lines_top", rename("quartz_block_lines_top", "quartz_pillar_top"))
	pipeline.AddPathContainsHandler("quartz_block_lines", rename("quartz_block_lines", "quartz_pillar"))
	pipeline.AddPathContainsHandler("quartz_block_chiseled", rename("quartz_block_chiseled", "chiseled_quartz_block"))

	// melon things and stems, pretty epic yo
	pipeline.AddPathContainsHandler("melon_stem_disconnected", rename("melon_stem_disconnected", "melon_stem"))
	pipeline.AddPathContainsHandler("melon_stem_connected", rename("melon_stem_connected", "attached_melon_stem"))
	pipeline.AddPathContainsHandler("pumpkin_stem_disconnected", rename("pumpkin_stem_disconnected", "pumpkin_stem"))
	pipeline.AddPathContainsHandler("pumpkin_stem_connected", rename("pumpkin_stem_connected", "attached_pumpkin_stem"))
	pipeline.AddPathContainsHandler("reeds", rename("reeds", "sugar_cane"))

	// cheating again cuz im lazy lol
	crops := [...]string{
		"wheat",
		"carrots",
		"potatoes",
		"nether_wart",
		"beetroots",
		"cocoa",
	}

	// lel
	for i := range crops {
		crop := crops[i]
		// just do for 10 stages, even though they only really go from 0 to 7
		for i2 := 0; i2 < 10; i2++ {
			num := strconv.FormatInt(int64(i2), 10)
			pipeline.AddPathContainsHandler(crop + "_stage_" + num, rename(crop + "_stage_" + num, crop + "_stage" + num))
		}
	}

	// logs, again, cheating repetition here
	logs := [...]string{
		"wood",
		"oak",
		"birch",
		"spruce",
		"jungle",
		"acacia",
		"big_oak",
		"dark_oak",
		"iron", // IRON IS WOOD? no, but it is a dor, just like i am lazy
	}

	for i := range logs {
		log := logs[i]

		pipeline.AddPathContainsHandler("log_" + log + "_top", rename("log_" + log + "_top", log + "_log_top"))
		pipeline.AddPathContainsHandler("planks_" + log, rename("planks_" + log, log + "_planks"))
		pipeline.AddPathContainsHandler("door_" + log + "_lower", rename("log_" + log + "_lower", log + "_door_bottom"))
		pipeline.AddPathContainsHandler("door_" + log + "_upper", rename("log_" + log + "_upper", log + "_door_top"))
		pipeline.AddPathContainsHandler("log_" + log, rename("log_" + log, log + "_log"))
		pipeline.AddPathContainsHandler("sapling_" + log, rename("sapling_" + log, log + "_sapling"))
		pipeline.AddPathContainsHandler("leaves_" + log, rename("leaves_" + log, log + "_leaves"))
		pipeline.AddPathContainsHandler("trapdoor_" + log, rename("trapdoor_" + log, log + "_trapdoor"))
	}

	// only replace the original trapdoor, but don't break new ones
	pipeline.AddPathContainsHandler("/trapdoor.", rename("/trapdoor." ,"oak_trapdoor"))

	// and now! back to boring conventions, thanks mojang for adding so many fucking flowers
	pipeline.AddPathContainsHandler("tallgrass", rename("tallgrass", "grass"))
	pipeline.AddPathContainsHandler("deadbush", rename("deadbush", "dead_bush"))
	pipeline.AddPathContainsHandler("flower_allium", rename("flower_allium", "allium"))
	pipeline.AddPathContainsHandler("flower_blue_orchid", rename("flower_blue_orchid", "blue_orchid"))
	pipeline.AddPathContainsHandler("flower_dandelion", rename("flower_dandelion", "dandelion"))
	pipeline.AddPathContainsHandler("flower_houstonia", rename("flower_houstonia", "azure_bluet"))
	pipeline.AddPathContainsHandler("flower_rose", rename("flower_rose", "poppy"))

	// let's just re use the fucking colors
	for i := range colors {
		color := colors[i]
		pipeline.AddPathContainsHandler("flower_tulip_" + color, rename("flower_tulip_" + color, color + "_tulip"))
	}

	// more double blocks, gotta love em
	pipeline.AddPathContainsHandler("double_plant_paeonia_bottom", rename("double_plant_paeonia_bottom", "peony_bottom"))
	pipeline.AddPathContainsHandler("double_plant_paeonia_top", rename("double_plant_paeonia_top", "peony_top"))
	pipeline.AddPathContainsHandler("double_plant_rose_bottom", rename("double_plant_rose_bottom", "rose_bush_bottom"))
	pipeline.AddPathContainsHandler("double_plant_rose_top", rename("double_plant_rose_top", "rose_bush_top"))

	// cute little sun flowers! i hope they burn
	pipeline.AddPathContainsHandler("double_plant_sunflower_bottom", rename("double_plant_sunflower_bottom", "sunflower_bottom"))
	pipeline.AddPathContainsHandler("double_plant_sunflower_top", rename("double_plant_sunflower_top", "sunflower_top"))
	pipeline.AddPathContainsHandler("double_plant_sunflower_back", rename("double_plant_sunflower_back", "sunflower_back"))
	pipeline.AddPathContainsHandler("double_plant_sunflower_front", rename("double_plant_sunflower_front", "sunflower_front"))

	// more double plants, because i was just wondering where my depression went
	pipeline.AddPathContainsHandler("double_plant_syringa_bottom", rename("double_plant_syringa_bottom", "lilac_bottom"))
	pipeline.AddPathContainsHandler("double_plant_syringa_top", rename("double_plant_syringa_top", "lilac_top"))
	pipeline.AddPathContainsHandler("double_plant_fern_top", rename("double_plant_fern_top", "large_fern_top"))
	pipeline.AddPathContainsHandler("double_plant_fern_bottom", rename("double_plant_fern_bottom", "large_fern_bottom"))
	pipeline.AddPathContainsHandler("double_plant_grass_top", rename("double_plant_grass_top", "tall_grass_top"))
	pipeline.AddPathContainsHandler("double_plant_grass_bottom", rename("double_plant_grass_bottom", "tall_grass_bottom"))

	// shrooms
	pipeline.AddPathContainsHandler("mushroom_block_skin_stem", rename("mushroom_block_skin_stem", "mushroom_stem"))
	pipeline.AddPathContainsHandler("mushroom_block_skin_brown", rename("mushroom_block_skin_brown", "brown_mushroom_block"))
	pipeline.AddPathContainsHandler("mushroom_block_skin_red", rename("mushroom_block_skin_red", "red_mushroom_block"))
	pipeline.AddPathContainsHandler("mushroom_red", rename("mushroom_red", "red_mushroom"))
	pipeline.AddPathContainsHandler("mushroom_brown", rename("mushroom_brown", "brown_mushroom"))

	// rails, yoink
	pipeline.AddPathContainsHandler("rail_normal_turned", rename("rail_normal_turned", "rail_corner"))
	pipeline.AddPathContainsHandler("rail_activator_powered", rename("rail_activator_powered", "activator_rail_on"))
	pipeline.AddPathContainsHandler("rail_activator", rename("rail_activator", "activator_rail"))
	pipeline.AddPathContainsHandler("rail_detector_powered", rename("rail_detector_powered", "detector_rail_on"))
	pipeline.AddPathContainsHandler("rail_detector", rename("rail_detector", "detector_rail"))
	pipeline.AddPathContainsHandler("rail_golden_powered", rename("rail_golden_powered", "powered_rail_on"))
	pipeline.AddPathContainsHandler("rail_golden", rename("rail_golden", "powered_rail"))
	pipeline.AddPathContainsHandler("rail_normal", rename("rail_normal", "rail"))

	// block states, fun, aye
	pipeline.AddPathContainsHandler("fire_layer_0", rename("fire_layer_0", "fire_0"))
	pipeline.AddPathContainsHandler("fire_layer_1", rename("fire_layer_1", "fire_1"))
	pipeline.AddPathContainsHandler("noteblock", rename("noteblock", "note_block"))
	pipeline.AddPathContainsHandler("slime.", rename("slime.", "slime_block"))
	pipeline.AddPathContainsHandler("trip_wire_hook", rename("trip_wire_hook", "tripwire_hook"))
	pipeline.AddPathContainsHandler("waterlily", rename("waterlily", "lily_pad"))
	pipeline.AddPathContainsHandler("ice_packed", rename("ice_packed", "packed_ice"))
	pipeline.AddPathContainsHandler("prismarine_dark", rename("prismarine_dark", "dark_prismarine"))
	pipeline.AddPathContainsHandler("prismarine_rough", rename("prismarine_rough", "prismarine"))
	pipeline.AddPathContainsHandler("trip_wire_source", rename("trip_wire_source", "trip_wire_hook"))
	pipeline.AddPathContainsHandler("hardened_clay", rename("hardened_clay", "terracotta"))
	pipeline.AddPathContainsHandler("sponge_wet", rename("sponge_wet", "wet_sponge"))
	pipeline.AddPathContainsHandler("anvil_top_damaged_2", rename("anvil_top_damaged_2", "damaged_anvil_top"))
	pipeline.AddPathContainsHandler("anvil_top_damaged_1", rename("anvil_top_damaged_1", "chipped_anvil_top"))
	pipeline.AddPathContainsHandler("anvil_top_damaged_0", rename("anvil_top_damaged_0", "anvil_top"))
	pipeline.AddPathContainsHandler("anvil_base", rename("anvil_base", "anvil"))
	pipeline.AddPathContainsHandler("piston_top_normal", rename("piston_top_normal", "piston_top"))
	pipeline.AddPathContainsHandler("endframe_top", rename("endframe_top", "end_portal_frame_top"))
	pipeline.AddPathContainsHandler("endframe_side", rename("endframe_side", "end_portal_frame_side"))
	pipeline.AddPathContainsHandler("endframe_eye", rename("endframe_eye", "end_portal_frame_eye"))
	pipeline.AddPathContainsHandler("end_bricks", rename("end_bricks", "end_stone_bricks"))
	pipeline.AddPathContainsHandler("pumpkin_face_off", rename("pumpkin_face_off", "carved_pumpkin"))
	pipeline.AddPathContainsHandler("pumpkin_face_on", rename("pumpkin_face_on", "jack_o_lantern"))
	pipeline.AddPathContainsHandler("web.", rename("web.", "cobweb"))
	pipeline.AddPathContainsHandler("comparator_off", rename("comparator_off", "comparator"))
	pipeline.AddPathContainsHandler("repeater_off", rename("repeater_off", "repeater"))
	pipeline.AddPathContainsHandler("redstone_torch_on", rename("redstone_torch_on", "redstone_torch"))
	pipeline.AddPathContainsHandler("torch_on", rename("torch_on", "torch"))
	pipeline.AddPathContainsHandler("observer_back_lit", rename("observer_back_lit", "observer_back_on"))
	pipeline.AddPathContainsHandler("dropper_front_horizontal", rename("dropper_front_horizontal", "dropper_front"))
	pipeline.AddPathContainsHandler("dispenser_front_horizontal", rename("dispenser_front_horizontal", "dispenser_front"))
	pipeline.AddPathContainsHandler("furnace_front_off", rename("furnace_front_off", "furnace_front"))
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
