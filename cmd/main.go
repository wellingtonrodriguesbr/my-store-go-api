package main

import (
	"my-store-api-go/db"
	"my-store-api-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	routes.ProductRoutes(server, dbConnection)

	server.Run(":8000")
}
