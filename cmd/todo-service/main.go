package main

import (
	"log/slog"
	"os"
)

func main() {
	// logger init
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
