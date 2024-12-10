package main

import (
	"example.com/events/db"
	"example.com/events/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() // sets up HTTP server

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
