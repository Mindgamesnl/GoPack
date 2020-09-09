package gopack

import (
	"github.com/sirupsen/logrus"
	"strings"
)

type ResourcePack struct {
	Meta           PackMcMeta
	Root           string
	FileCollection FileCollection
}

func FromZip(filename string) ResourcePack {
	logrus.Info(filename, " currently is ", readableSize(len(DataFromFile(filename))))
	root := TransPath("work/original/")

	files, zipErr := Unzip(filename, root)
	if zipErr != nil {
		panic(zipErr)
	}

	// make file collection
	collection := FileCollection{
		Root: root,
		AllFiles: files,
		NameToPath: make(map[string]Resource),
	}

	for i := range files {
		path := strings.Replace(files[i], root, "", -1)

		elements := strings.Split(path, TransPath("/"))
		name := elements[len(elements) - 1]

		// duplicate file names can occur within a pack, so we should handle it
		resource := Resource{
			name,
			name + "@" + RandomString(5),
			path,
			root + path,
		}

		collection.NameToPath[resource.UniqueName] = resource
	}

	mcData := PackMcMeta{}
	FileToStruct(root + "pack.mcmeta", &mcData)

	logrus.Info("Loaded pack: " + mcData.Pack.Description, " in format ", mcData.Pack.PackFormat)

	return ResourcePack{
		Meta: mcData,
		Root: root,
		FileCollection: collection,
	}
}