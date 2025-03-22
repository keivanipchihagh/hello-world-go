package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/hello-world-go/api/http"
)

func Start() {
	router := gin.Default()

	router.GET("tasks", http.GetTasks)
	router.GET("tasks/:id", http.GetTask)
	router.POST("tasks", http.AddTask)
	router.DELETE("tasks/:id", http.DeleteTask)
	router.PUT("tasks/:id", http.UpdateTask)

	router.Run()
}
