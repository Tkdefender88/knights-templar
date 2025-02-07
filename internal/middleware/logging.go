package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode  int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		wroteHeader:    false,
	}
}

func (rw *responseWriter) Status() int {
	return rw.statusCode
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	if rw.wroteHeader {
		return
	}

	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.wroteHeader = true
}

func LoggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			slog.Info("Incoming Request",
				"Method", r.Method,
				"Path", r.URL.EscapedPath(),
				"Status", wrapped.statusCode,
				"elapsed", fmt.Sprintf("%dus", time.Since(start).Microseconds()))
		})
	}
}
