package gopack

func Make113Pipeline() {
	pipeline := CreatePipeline("to 1.13", "work/113/")

	// apply updates from 1.12 to 1.13
	ApplyFlatteningUpdate(pipeline)

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(4))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
