package gopack

func Make111Pipeline() {
	pipeline := CreatePipeline("to 1.11 (remove secrets)", "work/111/", "1.11-1.12.zip")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(3))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	CompressAssets(pipeline)

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
