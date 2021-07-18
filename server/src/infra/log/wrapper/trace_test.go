package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTraces(t *testing.T) {
	trace := GetTraces(0, 0)
	assert.Equal(t, "trace.go", trace[0].File)
	assert.Contains(t, "github.com/satorunooshie/eat-with-dog-API/server/src/infra/log/wrapper", trace[0].Module)
	assert.Contains(t, "Trace", trace[0].Function)

	trace = GetTraces(0, 1)
	assert.Equal(t, "trace.go", trace[0].File)
	assert.Contains(t, "github.com/satorunooshie/eat-with-dog-API/server/src/infra/log/wrapper", trace[0].Module)
	assert.Contains(t, "GetTraces", trace[0].Function)

	trace = GetTraces(0, 2)
	assert.Equal(t, "trace_test.go", trace[0].File)
	assert.Contains(t, "github.com/satorunooshie/eat-with-dog-API/server/src/infra/log/wrapper", trace[0].Module)
	assert.Contains(t, "TestGetTraces", trace[0].Function)
}

func TestTrace(t *testing.T) {
	tc, _ := Trace(0)
	assert.Equal(t, "trace.go", tc.File)
	assert.Contains(t, "github.com/satorunooshie/eat-with-dog-API/server/src/infra/log/wrapper", tc.Module)
	assert.Contains(t, "Trace", tc.Function)

	tc, _ = Trace(1)
	assert.Equal(t, "trace_test.go", tc.File)
	assert.Contains(t, "github.com/satorunooshie/eat-with-dog-API/server/src/infra/log/wrapper", tc.Module)
	assert.Contains(t, "TestTrace", tc.Function)
}
