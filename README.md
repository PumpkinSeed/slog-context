# slog: Context handler

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)

A Context Handler for [slog](https://pkg.go.dev/log/slog) Go library.

## Install

```sh
go get github.com/PumpkinSeed/slog-context
```

## Usage

GoDoc: [https://pkg.go.dev/github.com/PumpkinSeed/slog-context](https://pkg.go.dev/github.com/PumpkinSeed/slog-context)

### Example

```go
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
}
```

### Output

```json
{"time":"2023-08-15T13:36:50.207418292+02:00","level":"ERROR","msg":"this is an error","number":12,"string":"data","struct":{"Number":42,"String":"struct_data"}}
```