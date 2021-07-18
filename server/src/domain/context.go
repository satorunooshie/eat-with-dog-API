package domain

import (
	"context"
)

type (
	readOnly   int
	masterOnly int
)

var (
	readOnlyKey   readOnly
	masterOnlyKey masterOnly
)

func WithReadOnly(ctx context.Context, isReadOnly bool) context.Context {
	return context.WithValue(ctx, readOnlyKey, isReadOnly)
}

func IsReadOnly(ctx context.Context) bool {
	isReadOnly, ok := ctx.Value(readOnlyKey).(bool)
	return ok && isReadOnly
}

func IsMasterOnly(ctx context.Context) bool {
	isMasterOnly, ok := ctx.Value(masterOnlyKey).(bool)
	return ok && isMasterOnly
}
