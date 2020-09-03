package gopack

import (
	"github.com/Mindgamesnl/gutils"
	"github.com/sirupsen/logrus"
)

func RegisterPipelines()  {
	logrus.Info("Starting pipelines")
	logrus.Info("Initializing pipelines took ", + gutils.TimeFunction(func() {
		Make111Pipeline()
		Make113Pipeline()
		Make115Pipeline()
		Make116Pipeline()
	}), "MS")
}
