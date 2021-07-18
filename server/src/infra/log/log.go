package log

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/satorunooshie/eat-with-dog-API/server/src/API"
	"github.com/satorunooshie/eat-with-dog-API/server/src/infra/log/wrapper"
)

const (
	tagPrefix = "app."
)

var sync bool

var logrusLevel = logrus.DebugLevel

func init() {
	logrus.SetFormatter(buildFormatter())
	logrus.SetLevel(logrusLevel)
}

// ログの書き込みを同期的に行う
func SetAsSync() {
	sync = true
}

func Dump(v interface{}) {
	wrapper.Dump(v)
}

func Print(v interface{}) {
	wrapper.Print(v)
}

func Header(v ...interface{}) {
	wrapper.Header(v...)
}

func Mark() {
	wrapper.Mark(3)
}

func Warn(ctx context.Context, err error, title string) {
	Packet{
		Title:     title,
		Err:       err,
		TraceData: fmt.Sprintf("%+v", err),
	}.Warn(ctx)
}

func Error(ctx context.Context, err error, title string) {
	Packet{
		Title:     title,
		Err:       err,
		TraceData: fmt.Sprintf("%+v", err),
	}.Error(ctx)
}

func Errorf(ctx context.Context, err error, title string, args ...interface{}) {
	Error(ctx, err, fmt.Sprintf(title, args...))
}

type Packet wrapper.Packet

func NewPacket() Packet {
	return Packet{}
}

// 普段は使わない
func (p Packet) Fatal(v ...interface{}) {
	p.setRequest(v)
	p.setDefaultTag("fatal")
	packet := wrapper.Packet(p)
	if sync {
		packet.Panic()
		return
	}
	sendAsync(packet.Panic)
}

func (p Packet) Warn(v ...interface{}) {
	p.setRequest(v)
	p.setDefaultTag("warn")
	packet := wrapper.Packet(p)
	if sync {
		packet.Warn()
		return
	}
	sendAsync(packet.Warn)
}

func (p Packet) Error(v ...interface{}) {
	p.setRequest(v)
	p.setDefaultTag("error")
	packet := wrapper.Packet(p)
	if sync {
		packet.Error()
		return
	}
	sendAsyncWithRecoverDebug(packet.Error)
}

func (p Packet) ErrorSync(v ...interface{}) {
	p.setRequest(v)
	p.setDefaultTag("error")
	packet := wrapper.Packet(p)
	packet.Error()
}

func (p Packet) Debug(v ...interface{}) {
	p.setRequest(v)
	p.setDefaultTag("debug")
	packet := wrapper.Packet(p)
	if sync {
		packet.Debug()
		return
	}
	sendAsync(packet.Debug)
}

func (p *Packet) AddData(v ...interface{}) *Packet {
	p.DataList = append(p.DataList, v...)
	return p
}

func (p *Packet) setRequest(v []interface{}) {
	if p.TraceData == nil {
		if p.TraceDepth == 0 {
			p.TraceDepth = 20
		}
		if p.TraceSkip == 0 {
			p.TraceSkip = 4
		}
		p.TraceData = wrapper.GetTraces(p.TraceDepth, p.TraceSkip)
	}

	if len(v) == 0 {
		return
	}
	if ctx, ok := v[0].(context.Context); ok {
		p.UserID = API.SafeUserIDString(ctx)
	}
}

func (p *Packet) setDefaultTag(tag string) {
	if p.Tag != "" {
		return
	}
	p.Tag = tagPrefix + tag
}

func sendAsync(fn func()) {
	go func() {
		defer func() {
			_ = recover()
		}()
		fn()
	}()
}

func sendAsyncWithRecoverDebug(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				Header("ERROR DEBUG")
				Dump(err)
			}
		}()
		fn()
	}()
}
