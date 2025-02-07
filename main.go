package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.com/Tkdefender88/coolestyp/internal/middleware"
)

func SetupLogger() {
	loggerOptions := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, loggerOptions))

	slog.SetDefault(logger)
}

func main() {
	logger := middleware.LoggingMiddleware()
	mux := http.NewServeMux()

	SetupLogger()

	assetsRoot := "assets/"
	assetsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(assetsRoot)))
	mux.Handle("/app/", noCacheMiddleware(assetsHandler))

	srv := http.Server{
		Addr:    ":42069",
		Handler: logger(mux),
	}

	shutdownChannel := make(chan os.Signal, 1)

	signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGKILL, syscall.SIGABRT)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			slog.Error("Error running server", "error", err)
		}
	}()

	slog.Info("Server started")

	<-shutdownChannel

	slog.Info("Shutting down")
	shutdownContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := srv.Shutdown(shutdownContext)
	if err != nil {
		slog.Error("Error shutting down the server", "error", err)
	}
}

func noCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
