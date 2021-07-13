package gopack

func Make117Pipeline() {
	pipeline := CreatePipeline("to 1.17 (remove secrets, flattening)", TransPath("work/117/"), "1.17.zip")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(7))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	MakeEverythingLowercase(pipeline)
	RenamePackFolders(pipeline)
	ConvertItems(pipeline, Get14Blocks())
	ConvertItems(pipeline, Get14Items())
	ConvertItems(pipeline, GetBlocks())
	ConvertItems(pipeline, GetItems())
	ForceContent(pipeline, GetForcedContent())
	MigrateLanguage(pipeline, GetLang())

	CompressResources(pipeline)

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
