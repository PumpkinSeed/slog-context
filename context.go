package slogcontext

import "context"

type contextKey string
type contextVal map[string]any

var (
	fields contextKey = "slog_fields"
)

func WithValue(parent context.Context, key string, val any) context.Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if v, ok := parent.Value(fields).(contextVal); ok {
		v[key] = val
		return context.WithValue(parent, fields, v)
	}
	v := contextVal{
		key: val,
	}
	return context.WithValue(parent, fields, v)
}
