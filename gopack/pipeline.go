package gopack

import (
	"strings"
)

type Pipeline struct {
	Name                      string
	Handlers                  []func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline)
	UntouchedResourceHandlers []func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline)
	ProcessedFileNames        []string
	OutFolder                 string
	FileCache                 map[string][]byte
	WriteQueue                map[string][]byte
	WrittenFiles              map[string][]byte
	OutputName                string
}

// create a pipeline with a name (which will also be the output directory/pack and version)
func CreatePipeline(name string, out string, outputName string) *Pipeline {
	return &Pipeline{
		Name:                      name,
		UntouchedResourceHandlers: []func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline){},
		Handlers:                  []func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline){},
		ProcessedFileNames:        []string{},
		OutFolder:                 out,
		FileCache:                 make(map[string][]byte),
		WriteQueue:                make(map[string][]byte),
		WrittenFiles:              make(map[string][]byte),
		OutputName:                outputName,
	}
}

// add a handler for every file, THIS DOES **NOT** COUNT AS A SPECIFIC RESOURCE HANDLING!
func (p *Pipeline) AddGlobalHandler(handler func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline)) {
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		handler(originalPack, resource, pipeline)
	})
}

// add a handler for all png files or something, so the filetype argument would be "png" without the dot
func (p *Pipeline) AddForFileType(filetype string, handler func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline)) {
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		if strings.HasSuffix(resource.Path, filetype) {
			p.ProcessedFileNames = append(p.ProcessedFileNames, resource.UniqueName)
			handler(originalPack, resource, pipeline)
		}
	})
}

// add a handler for files with a regex
func (p *Pipeline) AddPathContainsHandler(regex string, handler func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline)) {
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		if strings.Contains(resource.OsPath, regex) {
			handler(originalPack, resource, pipeline)
			p.ProcessedFileNames = append(p.ProcessedFileNames, resource.UniqueName)
		}
	})
}

// add a handler for files with a specific name
func (p *Pipeline) AddForFileName(fileName string, handler func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline)) {
	p.Handlers = append(p.Handlers, func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		if resource.ReadableName == fileName {
			handler(originalPack, resource, pipeline)
			p.ProcessedFileNames = append(p.ProcessedFileNames, resource.UniqueName)
		}
	})
}

// A handler for all files that haven't been touched yet
func (p *Pipeline) UnhandledFileHandler(handler func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline)) {
	p.UntouchedResourceHandlers = append(p.Handlers, func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		handler(originalPack, resource, pipeline)
	})
}

// Save all untouched files
func (p *Pipeline) SaveUntouched() {
	p.UnhandledFileHandler(func(originalPack ResourcePack, resource *Resource, pipeline *Pipeline) {
		// remove comments
		pipeline.SaveBytes(resource, resource.ContentAsBytes())
	})
}

func (p *Pipeline) Flush() int {
	a := p.WriteQueue
	for s := range a {
		p.actuallyWrite(s, a[s])
	}
	p.WriteQueue = make(map[string][]byte)
	return len(a)
}

func (p *Pipeline) actuallyWrite(targetFolder string, content []byte) {
	// create folder
	WriteToFile(targetFolder, content)
	p.WrittenFiles[targetFolder] = content
}

func (p *Pipeline) SaveBytes(resource *Resource, content []byte) {
	p.FileCache[resource.UniqueName] = content
	p.WriteQueue[p.OutFolder+resource.Path] = content
}

func (p *Pipeline) SaveFile(resource *Resource, content string) {
	p.SaveBytes(resource, []byte(content))
}
