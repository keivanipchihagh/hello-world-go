package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/hello-world-go/pkg/models"
)

var tasks = []models.Task{
	{
		Id:     "1",
		Title:  "1984",
		Author: "George Orwell",
	},
	{
		Id:     "2",
		Title:  "To Kill a Mockingbird",
		Author: "Harper Lee",
	},
	{
		Id:     "3",
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
	},
}

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	for _, task := range tasks {
		if task.Id == id {
			c.IndentedJSON(http.StatusOK, task)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
}

func AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Task Structure"})
		log.Println(err.Error())
		return
	}
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Not Found"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Task Structure"})
		log.Println(err.Error())
		return
	}

	for i, task := range tasks {
		if task.Id == id {
			tasks[i] = newTask
			c.IndentedJSON(http.StatusOK, newTask)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
}
