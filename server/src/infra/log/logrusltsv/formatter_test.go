package logrusltsv

import (
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type UserValue struct{}

func (*UserValue) String() string {
	return ".String() was\tused\n"
}

func TestFormatter(t *testing.T) {
	location, err := time.LoadLocation("Asia/Tokyo")
	assert.NoError(t, err)

	entry := &logrus.Entry{
		Data: logrus.Fields{
			"app":                  "test",
			"time":                 "field time",
			"msg":                  "field msg",
			"level":                "field level",
			"custom":               &UserValue{},
			"field\nwith\tspecial": "value\nwith\tspecial",
		},
		Time: time.Date(2006, 1, 2, 15, 05, 06, 07, location),
		Message: "test\a message\n",
		Level: logrus.ErrorLevel,
	}

	// Fields are sorted by alphabetical.
	// "time", "msg", "level", fields are prefixed with "field.".
	// Special characters are quoted.
	// Field values are converted to string with /String() if it has.
	// Line ends with "\n".
	f := &Formatter{}
	output, err := f.Format(entry)
	assert.NoError(t, err)
	assert.Equal(
		t,
		strings.Join(
			[]string{
				`time:2006-01-02T15:05:06+09:00`,
				`level:error`,
				`msg:test\a message\n`,
				`app:test`,
				`custom:.String() was\tused\n`,
				`field.level:field level`,
				`field.msg:field msg`,
				`field.time:field time`,
				`field\nwith\tspecial:value\nwith\tspecial`,
			},
			"\t",
		)+"\n",
		string(output),
	)
}
