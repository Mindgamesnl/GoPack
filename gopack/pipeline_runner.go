package gopack

import (
	"github.com/Mindgamesnl/GoPack/gopack/utils"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"time"
)

var Pipelines []*Pipeline

func AddPipeline(pipeline *Pipeline) {
	Pipelines = append(Pipelines, pipeline)
}

func RunPipelines(originalPack ResourcePack) {
	utils.ResetFileCache()
	for i := range Pipelines {
		pack := originalPack // copy the pack for every pipeline, to prevent destruction
		pipe := Pipelines[i] // copy the pipeline

		logrus.Info("Executing pipeline: " + pipe.Name)

		tasks := (len(pipe.Handlers) + len(pipe.UntouchedResourceHandlers)) * len(pack.FileCollection.NameToPath)

		bar := pb.Full.Start(tasks)
		bar.SetRefreshRate(time.Millisecond * 10)

		// go over all files yo, very epic
		for s := range pack.FileCollection.NameToPath {
			originalFile := pack.FileCollection.NameToPath[s]
			file := &originalFile
			// go over all handlers
			for pipeIncrementer := range pipe.Handlers {
				handler := pipe.Handlers[pipeIncrementer]
				handler(pack, file, pipe)
				bar.Increment()
			}

			for pipeIncrementer := range pipe.UntouchedResourceHandlers {
				handler := pipe.UntouchedResourceHandlers[pipeIncrementer]

				if !contains(pipe.ProcessedFileNames, file.UniqueName) {
					handler(pack, file, pipe)
				}
				bar.Increment()
			}
		}
		bar.Finish()
		logrus.Info("Flushing files..")
		pipe.FlushFiles()
		logrus.Info("Saved!")
	}
	logrus.Info("Done lol")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}