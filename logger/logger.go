package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func InitLogger(lvl string) error {
	if lvl == "local" {
		lvl = "debug"
	}

	logrusLvl, err := logrus.ParseLevel(lvl)
	if err != nil {
		return err
	}

	txtFormatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: false,
		ForceColors:            true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}

	logrus.SetLevel(logrusLvl)
	logrus.SetFormatter(txtFormatter)
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)

	return nil
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
