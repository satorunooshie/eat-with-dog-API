package log

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/satorunooshie/eat-with-dog-API/server/src/infra/log/logrusltsv"
	"github.com/satorunooshie/eat-with-dog-API/server/src/infra/log/wrapper"
)

func buildFormatter() logrus.Formatter {
	f := logrusltsv.New()
	f.AddIgnore("context")
	f.AddFilter("http_request", filterRequest)
	f.AddFilter("trace", filterTrace)
	return f
}

func filterRequest(v interface{}) interface{} {
	r, ok := v.(*http.Request)
	if !ok {
		return v
	}
	uri := FilterTokenFromURI(r.RequestURI)
	return fmt.Sprintf("%s -> %s %s%s", r.RemoteAddr, r.Method, r.Host, uri)
}

func filterTrace(v interface{}) interface{} {
	traces, ok := v.([]wrapper.StackTrace)
	if !ok {
		return v
	}

	max := 8
	size := len(traces)
	if len(traces) < max {
		max = size
	}

	list := make([]string, max)
	for i, v := range traces[:max] {
		list[i] = fmt.Sprintf("%s:L%d:%s", trimPath(v.Path), v.Line, v.Function)
	}
	return list
}

func trimPath(p string) string {
	const (
		path       = "/"
		pathLength = len(path)
	)

	idx := strings.Index(p, path)
	if idx == -1 {
		return p
	}
	return p[idx+pathLength:]
}
