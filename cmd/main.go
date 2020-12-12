package main

import (
	"core-backend/pkg/repository/sqlite"
	"net/http"

	"github.com/gin-gonic/gin"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hais",
		})
	})

	sqlite.MigrateDatabase()
	r.Run()
}
