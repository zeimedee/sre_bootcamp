package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_http_request_total",
			Help: "Total number of requests processed by sre_bootcamp",
		}, []string{"path", "status"})

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_http_request_error_total",
			Help: "Total number of errors requests processed by sre_bootcamp.",
		},
		[]string{"path", "status"})
)

var customRegistry = prometheus.NewRegistry()

func init() {
	customRegistry.MustRegister(RequestCount, ErrorCount)
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.HandlerFor(customRegistry, promhttp.HandlerOpts{})
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func TrackMetrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/metrics" {
			ctx.Next()
			return
		}
		path := ctx.FullPath()
		ctx.Next()
		status := ctx.Writer.Status()
		RequestCount.WithLabelValues(path, http.StatusText(status)).Inc()
		if status >= 400 {
			ErrorCount.WithLabelValues(path, http.StatusText(status)).Inc()
		}
	}
}
