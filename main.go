package main

import (
	"ethsyncer/cmd"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"time"
)

func init() {
	formatter := &log.TextFormatter{
		TimestampFormat: time.StampMilli,
		FullTimestamp:   true,
		PadLevelText:    true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			file := fmt.Sprintf(" [%s:%d]\t", path.Base(f.File), f.Line)
			return "", file
		},
	}
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(formatter)
}

func main() {
	cmd.Execute()
}
