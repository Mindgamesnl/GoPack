package gopack

func Make113Pipeline() {
	pipeline := CreatePipeline("to 1.13 (remove secrets, flattening)", "work/113/", "1.13-1.14.zip")


	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(4))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())


	MakeEverythingLowercase(pipeline)
	RenamePackFolders(pipeline)
	ConvertItems(pipeline, Get14Blocks())
	ConvertItems(pipeline, Get14Items())
	ForceContent(pipeline, GetForcedContent())
	MigrateLanguage(pipeline, GetLang14())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
