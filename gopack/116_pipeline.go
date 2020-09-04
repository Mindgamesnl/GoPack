package gopack

func Make116Pipeline() {
	pipeline := CreatePipeline("to 1.16 (remove secrets, flattening)", "work/116/", "1.16.zip")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(6))
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

	CompressAssets(pipeline)

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
