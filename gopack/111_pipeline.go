package gopack

func Make111Pipeline() {
	pipeline := CreatePipeline("to 1.11 (remove secrets)", "work/111/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(3))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
