package wrapper

import (
	"github.com/sirupsen/logrus"
)

const (
	LevelFatal = logrus.FatalLevel
	LevelPanic = logrus.PanicLevel
	LevelError = logrus.ErrorLevel
	LevelWarn  = logrus.WarnLevel
	LevelInfo  = logrus.InfoLevel
	LevelDebug = logrus.DebugLevel
)

var (
	LevelsPanic = []logrus.Level{
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
	LevelsError = append(LevelsPanic, logrus.ErrorLevel)
	LevelsWarn  = append(LevelsError, logrus.WarnLevel)
	LevelsInfo  = append(LevelsWarn, logrus.InfoLevel)
	LevelsDebug = append(LevelsInfo, logrus.DebugLevel)
)
