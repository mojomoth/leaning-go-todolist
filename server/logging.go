package server

import (
	"log"
	"net/http"
)

type statusCapturingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (s *statusCapturingResponseWriter) Write(b []byte) (int, error) {
	if s.status == 0 {
		s.WriteHeader(http.StatusOK)
	}

	return s.ResponseWriter.Write(b)
}

func (s *statusCapturingResponseWriter) WriteHeader(statusCode int) {
	s.status = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusCapturingWriter := &statusCapturingResponseWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(statusCapturingWriter, r)

		log.Printf("[%s] %s - %d\n", r.Method, r.RequestURI, statusCapturingWriter.status)
	})
}
