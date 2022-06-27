package prometheus

import (
	"bufio"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

type MetricsMiddleware struct {
	opsProcessed *prometheus.CounterVec
}

func NewMetricsMiddleware() *MetricsMiddleware {
	opsProcessed := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"method", "path", "statuscode"})
	return &MetricsMiddleware{
		opsProcessed: opsProcessed,
	}
}

// Metrics middleware to collect metrics from http requests
func (lm *MetricsMiddleware) Metrics() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {

		//wi := &responseWriterInterceptor{
		//	statusCode:     http.StatusOK,
		//	ResponseWriter: c.Writer,
		//}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// responseWriterInterceptor is a simple wrapper to intercept set data on a
// ResponseWriter.
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterInterceptor) Write(p []byte) (int, error) {
	return w.ResponseWriter.Write(p)
}

func (w *responseWriterInterceptor) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("type assertion failed http.ResponseWriter not a http.Hijacker")
	}
	return h.Hijack()
}

func (w *responseWriterInterceptor) Flush() {
	f, ok := w.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}

	f.Flush()
}
