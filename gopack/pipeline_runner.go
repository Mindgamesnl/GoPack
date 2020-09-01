package gopack

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
)

var Pipelines []*Pipeline

func AddPipeline(pipeline *Pipeline) {
	Pipelines = append(Pipelines, pipeline)
}

func RunPipelines(pack ResourcePack) {
	for i := range Pipelines {
		pipe := Pipelines[i]
		logrus.Info("Executing pipeline: " + pipe.Name)

		tasks := len(pack.FileCollection.AllFiles)

		bar := pb.StartNew(tasks)

		// go over all files yo, very epic
		for s := range pack.FileCollection.NameToPath {
			file := pack.FileCollection.NameToPath[s]
			// logrus.Info(len(pipe.Handlers))
			// go over all handlers
			for pipeIncrementer := range pipe.Handlers {
				handler := pipe.Handlers[pipeIncrementer]
				handler(pack, file, pipe)
			}

			for pipeIncrementer := range pipe.Handlers {
				handler := pipe.Handlers[pipeIncrementer]
				handler(pack, file, pipe)
			}

			for pipeIncrementer := range pipe.UntouchedResourceHandlers {
				handler := pipe.UntouchedResourceHandlers[pipeIncrementer]

				if !contains(pipe.ProcessedFileNames, file.UniqueName) {
					handler(pack, file, pipe)
				}
			}

			file.ContentAsString()

			bar.Increment()
		}

		bar.Finish()
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}