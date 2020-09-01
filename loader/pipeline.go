package loader

import "strings"

type Pipeline struct {
	Name                      string
	Handlers                  []func(originalPack ResourcePack, resource Resource, pipeline Pipeline)
	UntouchedResourceHandlers []func(originalPack ResourcePack, resource Resource, pipeline Pipeline)
	ProcessedFileNames        []string
}

// create a pipeline with a name (which will also be the output directory/pack and version)
func CreatePipeline(name string) Pipeline {
	return Pipeline{
		Name:                      name,
		UntouchedResourceHandlers: []func(originalPack ResourcePack, resource Resource, pipeline Pipeline){},
		Handlers:                  []func(originalPack ResourcePack, resource Resource, pipeline Pipeline){},
		ProcessedFileNames:        []string{},
	}
}

// add a handler for every file, THIS DOES **NOT** COUNT AS A SPECIFIC RESOURCE HANDLING!
func (p Pipeline) AddGlobalHandler(handler func(originalPack ResourcePack, resource Resource, pipeline Pipeline)) Pipeline {
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource Resource, pipeline Pipeline) {
		handler(originalPack, resource, pipeline)
	})
	return p
}

// add a handler for all png files or something, so the filetype argument would be "png" with or without the dot
func (p Pipeline) AddForFileType(filetype string, handler func(originalPack ResourcePack, resource Resource, pipeline Pipeline)) Pipeline {
	filetype = strings.Replace(filetype, ".", "", -1)
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource Resource, pipeline Pipeline) {
		if strings.Contains(resource.ReadableName, "."+filetype) {
			handler(originalPack, resource, pipeline)
			p.ProcessedFileNames = append(p.ProcessedFileNames, resource.UniqueName)
		}
	})
	return p
}

// add a handler for files with a specific name
func (p Pipeline) AddForFileName(fileName string, handler func(originalPack ResourcePack, resource Resource, pipeline Pipeline)) Pipeline {
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource Resource, pipeline Pipeline) {
		if resource.ReadableName == fileName {
			handler(originalPack, resource, pipeline)
			p.ProcessedFileNames = append(p.ProcessedFileNames, resource.UniqueName)
		}
	})
	return p
}

// A handler for all files that haven't been touched yet
func (p Pipeline) UnhandledFileHandler(handler func(originalPack ResourcePack, resource Resource, pipeline Pipeline)) Pipeline {
	p.UntouchedResourceHandlers = append(p.Handlers, func(originalPack ResourcePack, resource Resource, pipeline Pipeline) {
		handler(originalPack, resource, pipeline)
	})
	return p
}
