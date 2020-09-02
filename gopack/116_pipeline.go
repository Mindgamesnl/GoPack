package gopack

import "github.com/Mindgamesnl/GoPack/gopack/sources"

func Make116Pipeline() {
	pipeline := CreatePipeline("to 1.16 (remove secrets, flattening)", "work/116/")

	// apply updates from 1.12 to 1.15
	RenamePackFolders(pipeline)
	ConvertItems(pipeline, sources.Get14Blocks())
	ConvertItems(pipeline, sources.Get14Items())
	ConvertItems(pipeline, sources.GetBlocks())
	ConvertItems(pipeline, sources.GetItems())
	MakeEverythingLowercase(pipeline)

	// TODO: CONVERT LANGUAGE

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(6))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
