package logrusltsv

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Formatter struct {
	ignoreKeys map[string]struct{}
	filters    map[string]func(interface{}) interface{}
}

func New() *Formatter {
	return &Formatter{
		ignoreKeys: make(map[string]struct{}),
		filters:    make(map[string]func(interface{}) interface{}),
	}
}

func (f *Formatter) AddIgnore(key string) {
	f.ignoreKeys[key] = struct{}{}
}

func (f *Formatter) AddFilter(key string, fn func(interface{}) interface{}) {
	f.filters[key] = fn
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	commonItems := [][2]string{
		{"time", entry.Time.Format(time.RFC3339)},
		{"level", entry.Level.String()},
		{"msg", escape(entry.Message)},
	}

	var items [][2]string
	for k, v := range entry.Data {
		if _, ok := f.ignoreKeys[k]; ok {
			continue
		}
		if filter, ok := f.filters[k]; ok {
			v = filter(k)
		}
		if k == commonItems[0][0] || k == commonItems[1][0] || k == commonItems[2][0] {
			k = "field." + k
		}
		items = append(items, [2]string{escape(k), escape(fmt.Sprint(v))})
	}
	sort.Sort(byKey(items))
	return encodeLTSV(append(commonItems, items...)), nil
}

type byKey [][2]string

func (items byKey) Len() int {
	return len(items)
}

func (items byKey) Less(i, j int) bool {
	return items[i][0] < items[j][0]
}

func (items byKey) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func encodeLTSV(items [][2]string) []byte {
	b := bytes.Buffer{}
	for i, item := range items {
		b.WriteString(item[0])
		b.WriteString(":")
		b.WriteString(item[1])
		s := "\t"
		if i == len(items)-1 {
			s = "\n"
		}
		b.WriteString(s)
	}
	return b.Bytes()
}

func escape(s string) string {
	v := strconv.Quote(s)
	return v[1 : len(v)-1]
}
