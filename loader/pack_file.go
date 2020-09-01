package loader

import "github.com/Mindgamesnl/GoPack/utils"

type Resource struct {
	ReadableName string
	UniqueName   string
	Path         string
	OsPath       string
}

func (r Resource) ContentAsString() string {
	return utils.TextFromFile(r.OsPath)
}
