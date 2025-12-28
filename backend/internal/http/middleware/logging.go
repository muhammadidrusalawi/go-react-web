package middleware

import (
	"log"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		recorder := &statusRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		next.ServeHTTP(recorder, r)

		duration := time.Since(start).Microseconds()
		log.Printf("[HTTP] %s %s | %d | %dµs", r.Method, r.URL.Path, recorder.Status, duration)
	})
}
