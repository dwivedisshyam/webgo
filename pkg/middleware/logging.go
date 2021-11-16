package middeware

import (
	"net/http"
	"time"

	"github.com/dwivedisshyam/webgo/pkg/log"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *StatusResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func Logging(logger log.Logger) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			srw := &StatusResponseWriter{ResponseWriter: w}

			defer func(res *StatusResponseWriter) {
				responseTime := time.Since(start).Microseconds()

				if res.status >= http.StatusBadRequest {
					logger.Errorf("%s %s %d - %v", r.Method, r.URL.Path, res.status, responseTime)
				} else {
					logger.Infof("%s %s %d - %v", r.Method, r.URL.Path, res.status, responseTime)
				}
			}(srw)

			inner.ServeHTTP(srw, r)
		})
	}
}
