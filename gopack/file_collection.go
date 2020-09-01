package gopack

type FileCollection struct {
	Root string
	AllFiles []string
	NameToPath map[string]Resource
}
