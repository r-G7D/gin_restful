package main

import (
	"r-G7D/go_gin_restful/app"
	"r-G7D/go_gin_restful/handlers/driverHandler"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	app.DefDB()

	g.GET("/api/drivers", driverHandler.Index)
	g.GET("/api/drivers/:id", driverHandler.Show)
	g.POST("/api/drivers", driverHandler.Create)
	g.PUT("/api/drivers/:id", driverHandler.Update)
	g.DELETE("/api/drivers/:id", driverHandler.Delete)

	g.Run()
}
