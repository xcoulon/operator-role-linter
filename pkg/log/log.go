package log

import (
	"path/filepath"
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetReportCaller(true)
	formatter := &log.TextFormatter{
		TimestampFormat:        "2006-01-02 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// return the filename and the
			return "", "[" + filepath.Base(f.File) + ":" + strconv.Itoa(f.Line) + "]"
		},
	}
	log.SetFormatter(formatter)
	log.SetLevel(log.DebugLevel)
}
