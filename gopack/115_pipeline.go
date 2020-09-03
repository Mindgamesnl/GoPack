package gopack

import "github.com/Mindgamesnl/GoPack/gopack/sources"

func Make115Pipeline() {
	pipeline := CreatePipeline("to 1.15 (remove secrets, flattening)", "work/115/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(5))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	// apply updates from 1.12 to 1.15
	MakeEverythingLowercase(pipeline)
	RenamePackFolders(pipeline)
	ConvertItems(pipeline, sources.Get14Blocks())
	ConvertItems(pipeline, sources.Get14Items())
	ConvertItems(pipeline, sources.GetBlocks())
	ConvertItems(pipeline, sources.GetItems())
	ForceContent(pipeline, sources.GetForcedContent())
	MigrateLanguage(pipeline, sources.GetLang14())
	MigrateLanguage(pipeline, sources.GetLang())

	AddPipeline(pipeline)
}
