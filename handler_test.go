package slogcontext

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"strings"
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

func TestHandler_WithValue_ShouldNotAffectParent(t *testing.T) {
	stringBuilder := strings.Builder{}
	handler := NewHandler(slog.NewJSONHandler(&stringBuilder, nil))
	logger := slog.New(handler)
	ctx := WithValue(context.Background(), "k1", "v1")
	WithValue(ctx, "k2", "v2")

	logger.InfoContext(ctx, "this is an log")

	logMessage := jsonToMap(t, stringBuilder.String())
	_, ok := logMessage["k2"]
	if ok {
		t.Errorf("Log message shouldn't contain key added to different context")
	}
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

func jsonToMap(t *testing.T, jsonStr string) map[string]any {
	result := make(map[string]any)
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		t.Errorf("Failed serializing log message into a map: %v", err)
	}
	return result
}
