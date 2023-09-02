package slogcontext

import (
	"context"
	"io"
	"log/slog"
	"os"
	"testing"
)

type Struct struct {
	Number int64
	String string
}

func TestHandler(t *testing.T) {
	ctx := WithValue(context.Background(), "number", 12)
	ctx = WithValue(ctx, "string", "data")
	ctx = WithValue(ctx, "struct", Struct{
		Number: 42,
		String: "struct_data",
	})
	logger := slog.New(NewHandler(slog.NewJSONHandler(os.Stdout, nil)))

	logger.ErrorContext(ctx, "this is an error")
}

func TestHandlerConcurrent(t *testing.T) {
	ctx := WithValue(context.Background(), "number", 12)
	ctx = WithValue(ctx, "string", "data")
	ctx = WithValue(ctx, "struct", Struct{
		Number: 42,
		String: "struct_data",
	})
	logger := slog.New(NewHandler(slog.NewJSONHandler(io.Discard, nil)))

	for i := 0; i < 100; i++ {
		go logger.ErrorContext(ctx, "this is an error")
	}

}

// goos: linux
// goarch: amd64
// pkg: github.com/PumpkinSeed/slog-context
// cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
// BenchmarkHandler
// BenchmarkHandler-8   	  673506	      1527 ns/op	     144 B/op	       4 allocs/op
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
