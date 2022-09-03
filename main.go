package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Run('http://localhost:8080')
}