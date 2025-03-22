package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func helloWorldName(c *gin.Context) {
	name := c.Param("name")
	msg := fmt.Sprintf("Hello World, %v!", name)
	c.String(http.StatusOK, msg)
}

func main() {

	// Load .env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find .env file")
	}

	router := gin.Default()
	router.GET("hello-world", helloWorld)
	router.GET("hello-world/:name", helloWorldName)

	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	address := fmt.Sprintf("%v:%v", host, port)
	router.Run(address)
}
