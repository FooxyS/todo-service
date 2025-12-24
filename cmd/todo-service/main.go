package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/FooxyS/todo-service/internal/adapter/memory"
	"github.com/FooxyS/todo-service/internal/controller/rest"
	"github.com/FooxyS/todo-service/internal/usecase"
)

func main() {
	// logger init
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// DB
	db := memory.New()

	// usecase
	uc := usecase.New(db)

	// http handlers and router
	router := rest.NewRouter(uc)

	slog.Info("start server", "host", "localhost", "port", 8080)

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
