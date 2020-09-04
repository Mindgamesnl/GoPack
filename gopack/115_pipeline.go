package gopack

func Make115Pipeline() {
	pipeline := CreatePipeline("to 1.15 (remove secrets, flattening)", "work/115/", "1.15.zip")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(5))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	// apply updates from 1.12 to 1.15
	MakeEverythingLowercase(pipeline)
	RenamePackFolders(pipeline)
	ConvertItems(pipeline, Get14Blocks())
	ConvertItems(pipeline, Get14Items())
	ConvertItems(pipeline, GetBlocks())
	ConvertItems(pipeline, GetItems())
	ForceContent(pipeline, GetForcedContent())
	MigrateLanguage(pipeline, GetLang())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
