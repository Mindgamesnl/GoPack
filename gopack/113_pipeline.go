package gopack

import "github.com/Mindgamesnl/GoPack/gopack/sources"

func Make113Pipeline() {
	pipeline := CreatePipeline("to 1.13 (remove secrets, flattening)", "work/113/")


	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(4))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())


	MakeEverythingLowercase(pipeline)
	RenamePackFolders(pipeline)
	ConvertItems(pipeline, sources.Get14Blocks())
	ConvertItems(pipeline, sources.Get14Items())
	ForceContent(pipeline, sources.GetForcedContent())
	MigrateLanguage(pipeline, sources.GetLang14())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
