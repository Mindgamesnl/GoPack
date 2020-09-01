package gopack

func Make113Pipeline() {
	pipeline := CreatePipeline("to 1.13", "work/113/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(4))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	// apply updates from 1.12 to 1.13
	// migrations.MigrateFlattening(&pipeline)

	pipeline.SaveUntouched()
	AddPipeline(&pipeline)
}
