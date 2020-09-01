package loader

import (
	"github.com/Mindgamesnl/GoPack/packs"
	"github.com/Mindgamesnl/GoPack/utils"
	"github.com/sirupsen/logrus"
	"strings"
)

type ResourcePack struct {
	Meta packs.PackMcMeta
	Root string
	FileCollection FileCollection
}

func FromZip(filename string) ResourcePack {
	root := "work/original/"

	files, _ := utils.Unzip(filename, root)

	// make file collection
	collection := FileCollection{
		Root: root,
		AllFiles: files,
		NameToPath: make(map[string]Resource),
	}

	for i := range files {
		path := strings.Replace(files[i], root, "", -1)

		elements := strings.Split(path, "/")
		name := elements[len(elements) - 1]

		// duplicate file names can occur within a pack, so we should handle it
		resource := Resource{
			name,
			name + "@" + utils.RandomString(5),
			path,
			root + path,
		}

		collection.NameToPath[resource.UniqueName] = resource
	}

	mcData := packs.PackMcMeta{}
	utils.JsonToStruct(root + "pack.mcmeta", &mcData)

	logrus.Info("Loaded pack: " + mcData.Pack.Description, " in format ", mcData.Pack.PackFormat)

	return ResourcePack{
		Meta: mcData,
		Root: root,
		FileCollection: collection,
	}
}