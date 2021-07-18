package wrapper

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type StdoutMock struct {
	stdout *os.File
	stderr *os.File
	writer *os.File
	output chan string
}

func (m *StdoutMock) Set() {
	backupOut := os.Stdout
	backupErr := os.Stderr
	r, w, _ := os.Pipe()

	os.Stdout = w
	os.Stderr = w
	opChan := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		opChan <- buf.String()
	}()

	m.stdout = backupOut
	m.stderr = backupErr
	m.writer = w
	m.output = opChan
}

func (m *StdoutMock) Get() string {
	_ = m.writer.Close()
	os.Stdout = m.stdout
	os.Stderr = m.stderr
	return <-m.output
}

func TestNewLogField(t *testing.T) {
	f := newLogField([]interface{}{"test message"})
	assert.Equal(t, "test message", f["value"])

	f = newLogField([]interface{}{"title", "test message"})
	assert.Equal(t, "title", f["message"])
	assert.Equal(t, "test message", f["value"])
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Error("error message")
	output := buf.String()
	assert.Contains(t, output, `level=error`)
	assert.Contains(t, output, `value="error message"`)
}

func TestWarn(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Warn("warn message")
	output := buf.String()
	assert.Contains(t, output, `level=warn`)
	assert.Contains(t, output, `value="warn message"`)
}

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Info("info message")
	output := buf.String()
	assert.Contains(t, output, `level=info`)
	assert.Contains(t, output, `value="info message"`)
}

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Debug("debug message")
	output := buf.String()
	assert.Empty(t, output)

	logrus.SetLevel(logrus.DebugLevel)
	Debug("debug message")
	output = buf.String()
	assert.Contains(t, output, `level=debug`)
	assert.Contains(t, output, `value="debug message"`)
}

func TestDump(t *testing.T) {
	m := &StdoutMock{}
	m.Set()

	Dump("dump message")

	op := m.Get()
	assert.Contains(t, op, `(string) (len=12) "dump message"`)
}

func TestPrint(t *testing.T) {
	m := &StdoutMock{}
	m.Set()

	Print("print message")

	op := m.Get()
	assert.Contains(t, op, "print message")
}

func TestHeader(t *testing.T) {
	m := &StdoutMock{}
	m.Set()

	Header()
	op := m.Get()
	assert.Contains(t, op, "=============================================\n")

	m.Set()
	Header("header message")
	op = m.Get()
	assert.Contains(t, op, "===================== header message =====================\n")
}
