package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mmclinton/db"
	"github.com/mmclinton/routes"
)

func main() {
	db.InitializeDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
