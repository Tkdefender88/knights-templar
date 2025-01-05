package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mux := http.NewServeMux()

	assetsRoot := "assets/"
	assetsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(assetsRoot)))
	mux.Handle("/app/", noCacheMiddleware(assetsHandler))

	srv := http.Server{
		Addr:    ":42069",
		Handler: mux,
	}


	shutdownChannel := make(chan os.Signal)

	signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGKILL, syscall.SIGABRT)

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	slog.Info("Server started")

	<-shutdownChannel

	slog.Info("Shutting down")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func noCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
