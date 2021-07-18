package wrapper

import (
	"io"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: logrus.New(),
	}
}

func (l *Logger) SetOutput(w io.Writer) {
	l.Logger.Out = w
}

func (l *Logger) DisableOutput() {
	l.SetOutput(ioutil.Discard)
}

func (l *Logger) SetFormatter(f logrus.Formatter) {
	l.Logger.Formatter = f
}

func (l *Logger) SetLogLevel(level logrus.Level) {
	l.Logger.SetLevel(level)
}

func (l *Logger) NewPacket() *Packet {
	return &Packet{
		Logger: l,
	}
}
