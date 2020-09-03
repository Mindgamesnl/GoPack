package gopack

import (
	"github.com/sirupsen/logrus"
)

func RegisterPipelines()  {
	logrus.Info("Starting pipelines")
	Make111Pipeline()
	Make113Pipeline()
	Make115Pipeline()
	Make116Pipeline()
}
