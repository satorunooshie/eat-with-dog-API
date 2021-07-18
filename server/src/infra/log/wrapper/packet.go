package wrapper

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	defaultTraceDepth = 10
	defaultSkipDepth  = 4
)

type Packet struct {
	*Logger

	Title   string
	Data    interface{}
	Err     error
	Request *http.Request
	Context context.Context
	SQL     string
	Engine  string
	UserID  string
	Tag     string

	TraceDepth int
	TraceSkip  int
	TraceData  interface{}
	NoTrace    bool

	DataList []interface{}
}

func (p *Packet) SetLogger(l *Logger) *Packet {
	p.Logger = l
	return p
}

func (p *Packet) Fatal() {
	if p.Logger == nil {
		logrus.WithFields(p.createField()).Fatal(p.Title)
		return
	}
	p.Logger.WithFields(p.createField()).Fatal(p.Title)
}

func (p *Packet) Panic() {
	if p.Logger == nil {
		logrus.WithFields(p.createField()).Panic(p.Title)
		return
	}
	p.Logger.WithFields(p.createField()).Panic(p.Title)
}

func (p *Packet) Error() {
	if p.Logger == nil {
		logrus.WithFields(p.createField()).Error(p.Title)
		return
	}
	p.Logger.WithFields(p.createField()).Error(p.Title)
}

func (p *Packet) Warn() {
	if p.Logger == nil {
		logrus.WithFields(p.createField()).Warn(p.Title)
		return
	}
	p.Logger.WithFields(p.createField()).Warn(p.Title)
}

func (p *Packet) Info() {
	if p.Logger == nil {
		logrus.WithFields(p.createField()).Info(p.Title)
		return
	}
	p.Logger.WithFields(p.createField()).Info(p.Title)
}

func (p *Packet) Debug() {
	if p.Logger == nil {
		logrus.WithFields(p.createField()).Debug(p.Title)
		return
	}
	p.Logger.WithFields(p.createField()).Debug(p.Title)
}

func (p *Packet) AddData(d ...interface{}) *Packet {
	p.DataList = append(p.DataList, d...)
	return p
}

func (p *Packet) createField() logrus.Fields {
	f := logrus.Fields{}

	f["tag"] = p.Tag
	f["value"] = createLogValue(p.Data, p.DataList)

	if p.Request != nil {
		f["http_request"] = p.Request
	}

	if p.Context != nil {
		f["context"] = p.Context
	}

	if p.SQL != "" {
		f["query"] = p.SQL
		if p.Engine != "" {
			f["engine"] = p.Engine
		}
	}

	if p.UserID != "" {
		f["user_id"] = p.UserID
	}

	if p.Err != nil {
		f["error"] = p.Err
	}

	switch {
	case p.NoTrace:
		return f
	case p.TraceData != nil:
		f["trace"] = p.TraceData
		return f
	}

	if p.TraceDepth == 0 {
		p.TraceDepth = defaultTraceDepth
	}

	if p.TraceSkip == 0 {
		p.TraceSkip = defaultSkipDepth
	}

	if traces := GetTraces(p.TraceDepth, p.TraceSkip); len(traces) != 0 {
		f["trace"] = traces
	}

	return f
}

func createLogValue(v interface{}, extra []interface{}) interface{} {
	switch {
	case len(extra) == 0:
		return v
	case v == nil:
		return extra
	}
	var list []interface{}
	list = append(list, v)
	return append(list, extra...)
}
