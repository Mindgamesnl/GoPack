package gopack

func Make116Pipeline() {
	pipeline := CreatePipeline("to 1.16", "work/116/")

	// apply updates from 1.12 to 1.15
	ApplyFlatteningUpdate(pipeline)

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(6))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
