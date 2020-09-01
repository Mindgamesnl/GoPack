package loader

import (
	"os"
	"path/filepath"
	"strings"
)

type Pipeline struct {
	Name                      string
	Handlers                  []func(originalPack ResourcePack, resource Resource, pipeline *Pipeline)
	UntouchedResourceHandlers []func(originalPack ResourcePack, resource Resource, pipeline *Pipeline)
	ProcessedFileNames        []string
	OutFolder                 string
}

// create a pipeline with a name (which will also be the output directory/pack and version)
func CreatePipeline(name string, out string) Pipeline {
	return Pipeline{
		Name:                      name,
		UntouchedResourceHandlers: []func(originalPack ResourcePack, resource Resource, pipeline *Pipeline){},
		Handlers:                  []func(originalPack ResourcePack, resource Resource, pipeline *Pipeline){},
		ProcessedFileNames:        []string{},
		OutFolder:                 out,
	}
}

// add a handler for every file, THIS DOES **NOT** COUNT AS A SPECIFIC RESOURCE HANDLING!
func (p *Pipeline) AddGlobalHandler(handler func(originalPack ResourcePack, resource Resource, pipeline *Pipeline)) {
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource Resource, pipeline *Pipeline) {
		handler(originalPack, resource, pipeline)
	})
}

// add a handler for all png files or something, so the filetype argument would be "png" with or without the dot
func (p *Pipeline) AddForFileType(filetype string, handler func(originalPack ResourcePack, resource Resource, pipeline *Pipeline)) {
	filetype = strings.Replace(filetype, ".", "", -1)
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource Resource, pipeline *Pipeline) {
		if strings.Contains(resource.ReadableName, "."+filetype) {
			handler(originalPack, resource, pipeline)
			p.ProcessedFileNames = append(p.ProcessedFileNames, resource.UniqueName)
		}
	})
}

// add a handler for files with a specific name
func (p *Pipeline) AddForFileName(fileName string, handler func(originalPack ResourcePack, resource Resource, pipeline *Pipeline)) {
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource Resource, pipeline *Pipeline) {
		if resource.ReadableName == fileName {
			handler(originalPack, resource, pipeline)
			p.ProcessedFileNames = append(p.ProcessedFileNames, resource.UniqueName)
		}
	})
}

// A handler for all files that haven't been touched yet
func (p *Pipeline) UnhandledFileHandler(handler func(originalPack ResourcePack, resource Resource, pipeline *Pipeline)) {
	p.UntouchedResourceHandlers = append(p.Handlers, func(originalPack ResourcePack, resource Resource, pipeline *Pipeline) {
		handler(originalPack, resource, pipeline)
	})
}

// Save all untouched files
func (p *Pipeline) SaveUntouched() {
	p.UnhandledFileHandler(func(originalPack ResourcePack, resource Resource, pipeline *Pipeline) {
		// remove comments
		pipeline.SaveBytes(resource, resource.ContentAsBytes())
	})
}

func (p Pipeline) SaveBytes(resource Resource, content []byte) {
	// create folder
	fa := os.MkdirAll(filepath.Dir(p.OutFolder+resource.Path), os.ModePerm)
	if fa != nil {
		panic(fa)
	}

	f, err := os.Create(p.OutFolder + resource.Path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err2 := f.Write(content)

	if err2 != nil {
		panic(err2)
	}
}


func (p Pipeline) SaveFile(resource Resource, content string) {
	p.SaveBytes(resource, []byte(content))
}
