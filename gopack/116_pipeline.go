package gopack

import "github.com/Mindgamesnl/GoPack/gopack/sources"

func Make116Pipeline() {
	pipeline := CreatePipeline("to 1.16 (remove secrets, flattening)", "work/116/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(6))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	MakeEverythingLowercase(pipeline)
	RenamePackFolders(pipeline)
	ConvertItems(pipeline, sources.Get14Blocks())
	ConvertItems(pipeline, sources.Get14Items())
	ConvertItems(pipeline, sources.GetBlocks())
	ConvertItems(pipeline, sources.GetItems())
	ForceContent(pipeline, sources.GetForcedContent())
	MigrateLanguage(pipeline, sources.GetLang14())
	MigrateLanguage(pipeline, sources.GetLang())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
