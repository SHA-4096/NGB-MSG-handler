package util

import (
	config "NGB-MSG-handler/internal/conf"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func LogUtilInit() {
	path := config.Config.LogConfig.LogPath
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(config.Config.LogConfig.MaxAge*60)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(config.Config.LogConfig.RotateTime*60)*time.Second),
	)
	multiWriter := io.MultiWriter(os.Stdout, writer)
	logrus.SetOutput(multiWriter)
}

func MakeInfoLog(msg string) {
	logrus.Info(msg)
}

func MakeWarnLog(msg string) {
	logrus.Warn(msg)
}

func MakeErrorLog(msg string) {
	logrus.Error(msg)
}
