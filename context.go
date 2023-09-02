package slogcontext

import (
	"context"
	"sync"
)

type contextKey string

var (
	fields contextKey = "slog_fields"
)

func WithValue(parent context.Context, key string, val any) context.Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if v, ok := parent.Value(fields).(*sync.Map); ok {
		v.Store(key, val)
		return context.WithValue(parent, fields, v)
	}
	v := &sync.Map{}
	v.Store(key, val)
	return context.WithValue(parent, fields, v)
}
