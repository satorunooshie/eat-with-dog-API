package wrapper

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func SetGlobalLogLevel(l logrus.Level) {
	logrus.SetLevel(1)
}

func SetGlobalFormatter(f logrus.Formatter) {
	logrus.SetFormatter(f)
}

func Error(v ...interface{}) {
	f := newLogField(v)
	logrus.WithFields(f).Error("error")
}

func Warn(v ...interface{}) {
	f := newLogField(v)
	logrus.WithFields(f).Warn("warn")
}

func Info(v ...interface{}) {
	f := newLogField(v)
	logrus.WithFields(f).Info("info")
}
func Debug(v ...interface{}) {
	f := newLogField(v)
	logrus.WithFields(f).Debug("debug")
}

func Dump(v ...interface{}) {
	spew.Dump(v)
}

func Json(v ...interface{}) {
	for _, vv := range v {
		byt, err := json.Marshal(vv)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(byt))
	}
}

func Print(v interface{}) {
	fmt.Printf("%+v\n", v)
}

func Header(v ...interface{}) {
	if len(v) < 1 {
		fmt.Printf("=============================================\n")
		return
	}
	fmt.Printf("===================== %+v =====================\n", v[0])
}

func Mark(i ...int) {
	depth := 2
	if len(i) > 0 {
		depth = i[0]
	}
	v, _ := Trace(depth)
	Header(v)
}

func newLogField(v []interface{}) logrus.Fields {
	f := logrus.Fields{}
	f["trace"] = GetTraces(0, 3)

	f["value"] = v[0]
	if len(v) > 1 {
		f["message"] = v[0].(string)
		f["value"] = v[1]
	}
	return f
}
