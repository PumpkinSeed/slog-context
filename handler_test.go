package slogcontext

import (
	"context"
	"io"
	"log/slog"
	"testing"
)

type Struct struct {
	Number int64
	String string
}

// goos: linux
// goarch: amd64
// pkg: github.com/PumpkinSeed/slog-context
// cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
// BenchmarkHandler
// BenchmarkHandler-8   	  864600	      1358 ns/op	     128 B/op	       3 allocs/op
func BenchmarkHandler(b *testing.B) {
	b.ReportAllocs()
	ctx := WithValue(context.Background(), "number", 12)
	ctx = WithValue(ctx, "string", "data")
	ctx = WithValue(ctx, "struct", Struct{
		Number: 42,
		String: "struct_data",
	})
	logger := slog.New(NewHandler(slog.NewJSONHandler(io.Discard, nil)))

	for i := 0; i < b.N; i++ {
		logger.ErrorContext(ctx, "this is an error")
	}
}
