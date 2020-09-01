package gopack

type Resource struct {
	ReadableName string
	UniqueName   string
	Path         string
	OsPath       string
}

func (r Resource) ContentAsString() string {
	return TextFromFile(r.OsPath)
}

func (r Resource) ContentAsBytes() []byte {
	return DataFromFile(r.OsPath)
}

func (r Resource) GetPipelineString(pipeline *Pipeline) string {
	return string(r.GetPipelineContent(pipeline))
}


func (r Resource) GetPipelineContent(pipeline *Pipeline) []byte {
	if val, ok := pipeline.FileCache[r.UniqueName]; ok {
		return val
	}

	content := r.ContentAsBytes()
	pipeline.FileCache[r.UniqueName] = content
	return content
}
