package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/PumpkinSeed/slog-context"
)

type Struct struct {
	Number int64
	String string
}

func main() {
	// Add slogcontext.Handler to slog
	slog.SetDefault(slog.New(slogcontext.NewHandler(slog.NewJSONHandler(os.Stdout, nil))))

	// Add values to context
	ctx := slogcontext.WithValue(context.Background(), "number", 12)
	ctx = slogcontext.WithValue(ctx, "string", "data")
	ctx = slogcontext.WithValue(ctx, "struct", Struct{
		Number: 42,
		String: "struct_data",
	})

	// Perform error log
	slog.ErrorContext(ctx, "this is an error")
	// out: {"time":"2023-08-15T13:36:50.207418292+02:00","level":"ERROR","msg":"this is an error","number":12,"string":"data","struct":{"Number":42,"String":"struct_data"}}
}
