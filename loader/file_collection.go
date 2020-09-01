package loader

type FileCollection struct {
	Root string
	AllFiles []string
	NameToPath map[string]string
}
