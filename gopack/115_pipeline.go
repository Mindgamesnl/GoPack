package gopack

func Make115Pipeline() {
	pipeline := CreatePipeline("to 1.15 (remove secrets, flattening)", "work/115/")

	// apply updates from 1.12 to 1.15
	ApplyFlatteningUpdate(pipeline)

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(5))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
