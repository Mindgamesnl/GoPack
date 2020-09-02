package gopack

import (
	"github.com/Mindgamesnl/GoPack/gopack/utils"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"strings"
)

var jsonCache = make(map[string]gjson.Result)
var textureNameMap = make(map[string]string)

func loadParsedJson(osPath string, resource *Resource, pipeline *Pipeline) gjson.Result {
	result, found := jsonCache[osPath]

	if found {
		return result
	}

	values := gjson.Parse(resource.GetPipelineString(pipeline))
	jsonCache[osPath] = values
	return values
}

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

	// do file mappings
	ApplyFlatteningMapping(pipeline)


	// Force all file names to be LOWERCASE, regardless of their type or path
	// more info can be found in the last json migration of this update
	pipeline.AddGlobalHandler(func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		resource.Path = strings.ToLower(resource.Path)
		resource.ReadableName = strings.ToLower(resource.ReadableName)
		resource.UniqueName = strings.ToLower(resource.UniqueName)
		pipeline.SaveBytes(resource, resource.GetPipelineContent(pipeline))
	})

	// process all json files
	pipeline.AddForFileType("json", func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// search json
		parsed := loadParsedJson(resource.OsPath, resource, pipeline)
		scan := utils.FindJsonKeys(parsed, resource.OsPath)

		updatedJson := resource.GetPipelineString(pipeline)
		updated := false

		for i := range scan {
			key := scan[i]

			// handler for texture files
			// rename texture references in json
			// but rebuild the json, dont actually set it, since keys can be
			// a number like fuck.0.you.all.3.bloody.4.cunts
			// and still will then explode for referencing
			// fuck.anotherkey.you
			// since it things that the second level is an array, while in fact, it's not
			if strings.Contains(key, "textures") {
				value := parsed.Get(key)
				// replace references to blocks/ and items/
				asString := value.Str

				asString = strings.Replace(asString, "items/", "item/", -1)
				asString = strings.Replace(asString, "blocks/", "block/", -1)
				asString = strings.ToLower(asString)

				for s := range textureNameMap {
					asString = strings.Replace(asString, s, textureNameMap[s], -1)
				}

				// one up
				parent := key
				parentParts := strings.Split(parent, ".")
				parent = strings.Replace(parent, "."+parentParts[len(parentParts)-1], "", -1)

				topLevel := parentParts[len(parentParts)-1]

				updatedMap := make(map[string]string)
				props := gjson.Get(updatedJson, parent).Map()

				for s := range props {
					v := props[s].Str
					if v != "" {
						updatedMap[s] = v
					}
				}

				updatedMap[topLevel] = asString

				var err error
				updatedJson, err = sjson.Set(updatedJson, parent, updatedMap)
				if err != nil {
					logrus.Info("item ", resource.GetPipelineString(pipeline))
					logrus.Info("key " + key)
					logrus.Info("Note: in texture")
					panic(err)
				}
				updated = true
			} else

			// recursively find model references and force them to be lower case
			// since minecraft >1.12.2 doesn't allow for uppercase characters
			// the files already got renamed in a few migrations before, so
			// this can be done safely
			// you might ask, why?
			// because minecraft, fuck you.
			if strings.Contains(key, "model") {
				value := parsed.Get(key)
				if getTypeCache(key, value) == gjson.String {
					asString := value.Str

					asString = strings.ToLower(asString)

					// and translate it, again
					for s := range textureNameMap {
						asString = strings.Replace(asString, s, textureNameMap[s], -1)
					}

					var err error
					updatedJson, err = sjson.Set(updatedJson, key, asString)
					if err != nil {
						logrus.Info("item ", resource.GetPipelineString(pipeline))
						logrus.Info("key " + key)
						logrus.Info("Note: in model")
						panic(err)
					}
					updated = true
				}
			}
		}

		if updated {
			if len(updatedJson) < 5 {
				logrus.Info("Writing small file, ", resource.Path)
			}
			pipeline.SaveFile(resource, updatedJson)
		}
	})

	// Save all untouched files
	// Shouldn't be any, since we rename like every fucking file
	// but i'd rather be safe then sorry and calling it anyway shouldn't be expensive
	// since it just does checks for files to write, and only does IO when there's something
	// in the pipeline, which again, shouldn't be anything in the first place.
	pipeline.SaveUntouched()
}

var typeCache = make(map[string]gjson.Type)

// getting types is slow as shit
// so caching it helps a lot
func getTypeCache(key string, res gjson.Result) gjson.Type {
	fc, found := typeCache[key]
	if found {
		return fc
	}
	ft := res.Type
	typeCache[key] = ft
	return ft
}

func rename(from string, to string) func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
	// build the map of old name > new name while we're at it
	textureNameMap[from] = to

	return func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// set and apply new name
		if strings.Contains(resource.Path, "textures/") || strings.Contains(resource.Path, "models/") {
			resource.Path = strings.Replace(resource.Path, from, to, 1)
			resource.ReadableName = strings.Replace(resource.ReadableName, from, to, 1)
			resource.UniqueName = strings.Replace(resource.UniqueName, from, to, 1)
			pipeline.SaveBytes(resource, resource.GetPipelineContent(pipeline))
		}
	}
}
