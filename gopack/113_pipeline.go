package gopack

import "github.com/Mindgamesnl/GoPack/gopack/sources"

func Make113Pipeline() {
	pipeline := CreatePipeline("to 1.13 (remove secrets, flattening)", "work/113/")

	// apply updates from 1.12 to 1.13
	RenamePackFolders(pipeline)
	MakeEverythingLowercase(pipeline)
	ConvertItems(pipeline, sources.Get14Blocks())
	ConvertItems(pipeline, sources.Get14Items())

	// TODO: CONVERT LANGUAGE

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(4))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
