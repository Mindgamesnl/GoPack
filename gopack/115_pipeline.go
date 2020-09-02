package gopack

import "github.com/Mindgamesnl/GoPack/gopack/sources"

func Make115Pipeline() {
	pipeline := CreatePipeline("to 1.15 (remove secrets, flattening)", "work/115/")

	// apply updates from 1.12 to 1.15
	RenamePackFolders(pipeline)
	ConvertItems(pipeline, sources.Get14Blocks())
	ConvertItems(pipeline, sources.Get14Items())
	MakeEverythingLowercase(pipeline)

	// TODO: CONVERT LANGUAGE

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(5))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
