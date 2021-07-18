package wrapper

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPacketError(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	p := &Packet{
		Title: "title",
		Data:  999,
		Tag:   "tag",
	}
	p.Error()
	output := buf.String()

	assert.Contains(t, output, `level=error`)
	assert.Contains(t, output, `msg=title`)
	assert.Contains(t, output, `tag=tag`)
	assert.Contains(t, output, `value=999`)
	trace := GetTraces(0, 2)
	assert.Contains(t, output, trace[0].Function)
}

func TestPacketInfo(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	p := &Packet{
		Title: "title",
		Data:  999,
		Tag:   "tag",
	}
	p.Info()
	output := buf.String()

	assert.Contains(t, output, `level=info`)
	assert.Contains(t, output, `msg=title`)
	assert.Contains(t, output, `tag=tag`)
	assert.Contains(t, output, `value=999`)
	trace := GetTraces(0, 2)
	assert.Contains(t, output, trace[0].Function)
}

func TestPacketCreateField(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	p := &Packet{
		Title: "title",
		Data:  999,
		Tag:   "tag",
	}
	f := p.createField()

	assert.Equal(t, p.Data, f["value"])
	assert.Equal(t, p.Tag, f["tag"])

	p.AddData("foo", "bar", 111)
	f = p.createField()
	assert.Equal(t, []interface{}{999, "foo", "bar", 111}, f["value"])
}

func TestPacketNoTrace(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	p := &Packet{
		Title:   "title",
		Data:    999,
		Tag:     "tag",
		NoTrace: true,
	}
	p.Error()
	output := buf.String()

	assert.Contains(t, output, `level=error`)
	assert.Contains(t, output, `msg=title`)
	assert.Contains(t, output, `tag=tag`)
	assert.Contains(t, output, `value=999`)
	trace := GetTraces(0, 2)
	assert.NotContains(t, output, trace[0].Function)
}
