package gopack

import (
	"archive/zip"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var Pipelines []*Pipeline

func AddPipeline(pipeline *Pipeline) {
	Pipelines = append(Pipelines, pipeline)
}

func RunPipelines(originalPack ResourcePack) {
	ResetFileCache()
	for i := range Pipelines {
		pack := originalPack // copy the pack for every pipeline, to prevent destruction
		pipe := Pipelines[i] // copy the pipeline

		tasks := (len(pipe.Handlers) + len(pipe.UntouchedResourceHandlers)) * len(pack.FileCollection.NameToPath)
		ntp := pack.FileCollection.NameToPath
		loadedCount := 0
		for s := range ntp {
			originalFile := ntp[s]
			file := &originalFile
			if strings.Contains(file.Path, ".") {
				loadedCount++
				file.GetPipelineContent(pipe)
			}
		}

		tmpl := `{{ red "` + pipe.Name + `:" }} {{percent .}} {{ bar . "[" "=" ">" " "}} {{speed . | rndcolor }} {{rtime . "ETA %s"}}`
		bar := pb.ProgressBarTemplate(tmpl).Start(tasks)
		bar.SetRefreshRate(time.Millisecond * 10)

		// go over all files yo, very epic
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
		}

		bar.Finish()

		pipe.Flush(len(pipe.FileCache))
		out := pipe.WrittenFiles

		for s := range out {
			if !fileExists(s) {
				logrus.Info("File doesn't really exist, " + s)
				panic("File doesn't really exist, " + s)
			} else {
				itsOwnData := DataFromFile(s)
				if string(out[s]) != string(itsOwnData) {
					logrus.Info("expected ", string(out[s]))
					logrus.Info("has ", string(itsOwnData))
					panic(s + " does not equal")
				}
			}
		}

		zipName := "out/" + pipe.OutputName

		fer := os.MkdirAll(filepath.Dir(zipName), os.ModePerm)
		if fer != nil {
			panic(fer)
		}

		// Get a Buffer to Write To
		outFile, err := os.Create(zipName)
		if err != nil {
			panic(err)
		}
		defer outFile.Close()

		// Create a new zip archive.
		w := zip.NewWriter(outFile)

		// Add some files to the archive.
		addFiles(w, pipe.OutFolder, "")

		if err != nil {
			panic(err)
		}

		err = w.Close()
		if err != nil {
			panic(err)
		}

		// read the zip and find the file size
		logrus.Info(zipName, " turned out to be ", readableSize(len(DataFromFile(zipName))))
	}
	logrus.Info("Finished pipeline, cleaning up..")

	// os.RemoveAll("work/")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
