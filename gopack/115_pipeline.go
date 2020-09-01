package gopack

func Make115Pipeline() {
	pipeline := CreatePipeline("to 1.15", "work/115/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(5))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	// apply updates from 1.12 to 1.15
	// migrations.MigrateFlattening(&pipeline)

	pipeline.SaveUntouched()
	AddPipeline(&pipeline)
}
