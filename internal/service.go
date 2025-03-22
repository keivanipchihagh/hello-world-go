package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/hello-world-go/api/http"
	"github.com/keivanipchihagh/hello-world-go/internal/config"
	"github.com/keivanipchihagh/hello-world-go/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start() {

	config := config.NewConfig()
	router := gin.Default()

	// Register Prometheus middleware
	router.Use(metrics.PrometheusMetrics())

	// Register API routes
	router.GET("tasks", http.GetTasks)
	router.GET("tasks/:id", http.GetTask)
	router.POST("tasks", http.AddTask)
	router.DELETE("tasks/:id", http.DeleteTask)
	router.PUT("tasks/:id", http.UpdateTask)

	// Expose Prometheus metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	router.Run(address)
}
