package gopack

import (
	"github.com/Mindgamesnl/GoPack/gopack/utils"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
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
		ntp := pack.FileCollection.NameToPath
		for s := range ntp {
			originalFile := ntp[s]

			file := &originalFile

			if strings.Contains(file.OsPath, ".") {
				// go over all handlers
				hdlr := pipe.Handlers
				for pipeIncrementer := range hdlr {
					handler := hdlr[pipeIncrementer]
					handler(pack, file, pipe)
					bar.Increment()
				}

				utv := pipe.UntouchedResourceHandlers
				for pipeIncrementer := range utv {
					handler := utv[pipeIncrementer]

					if !contains(pipe.ProcessedFileNames, file.UniqueName) {
						handler(pack, file, pipe)
					}
					bar.Increment()
				}
			} else {
				bar.Increment()
				bar.Increment()
			}

			ntp[s] = *file
			pack.FileCollection.NameToPath[s] = *file

			pipe.Flush()
			time.Sleep(1 * time.Millisecond)
		}
		pipe.Flush()
		bar.Finish()
		logrus.Info("Converting done. Validating written files...")
		out := pipe.WrittenFiles
		for s := range out {
			if !fileExists(s) {
				logrus.Info("File doesn't really exist, " + s)
				panic("File doesn't really exist, " + s)
			} else {
				itsOwnData := utils.DataFromFile(s)
				if string(out[s]) != string(itsOwnData) {
					panic(s + " does not equal")
				}
			}
		}

		logrus.Info("Files seem OK")
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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}