package gopack

import (
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"strings"
)

var jsonCache = make(map[string]gjson.Result)

func loadParsedJson(osPath string, resource *Resource, pipeline *Pipeline) gjson.Result {
	values := gjson.Parse(resource.GetPipelineString(pipeline))

	return values
}

func ConvertItems(pipeline *Pipeline, set map[string]string) {
	for s := range set {
		oldName := s
		updatedName := set[s]

		converter := func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
			// don't save original
			if strings.Contains(resource.Path, "/" + oldName + ".") {
				ogContent := resource.GetPipelineContent(pipeline)
				delete(pipeline.WriteQueue, pipeline.OutFolder+resource.Path)

				resource.Path = strings.Replace(resource.Path, oldName, updatedName, 1)
				resource.ReadableName = strings.Replace(resource.ReadableName, oldName, updatedName, 1)
				resource.UniqueName = strings.Replace(resource.UniqueName, oldName, updatedName, 1)

				pipeline.SaveBytes(resource, ogContent)
			}
		}

		// rename these files
		pipeline.AddGlobalHandler(converter)
	}

	// references
	pipeline.AddForFileType("json", func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// search json
		parsed := loadParsedJson(resource.OsPath, resource, pipeline)
		scan := FindJsonKeys(parsed, resource.OsPath)

		updatedJson := resource.GetPipelineString(pipeline)
		updated := false

		for i := range scan {
			key := scan[i]

			// handler for texture files
			// rename texture references in json
			// but rebuild the json, dont actually set it, since keys can be
			// a number like fuck.0.you.all.3.bloody.4.idiot
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

				for s := range set {
					if strings.HasSuffix(asString, "/" + s) && !strings.Contains(asString, "custom/") {
						asString = strings.Replace(asString, s, set[s], -1)
					}
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
					for s := range set {
						if strings.HasSuffix(asString, "/" + s) && !strings.Contains(asString, "custom/") {
							asString = strings.Replace(asString, s, set[s], -1)
						}
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
			delete(jsonCache, resource.OsPath)
		}
	})
}
