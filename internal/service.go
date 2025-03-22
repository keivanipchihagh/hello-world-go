package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/hello-world-go/api/http"
	"github.com/keivanipchihagh/hello-world-go/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start() {
	router := gin.Default()

	router.Use(metrics.PrometheusMetrics())

	// Register your API routes
	router.GET("tasks", http.GetTasks)
	router.GET("tasks/:id", http.GetTask)
	router.POST("tasks", http.AddTask)
	router.DELETE("tasks/:id", http.DeleteTask)
	router.PUT("tasks/:id", http.UpdateTask)

	// Expose Prometheus metrics on /metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Run()
}
