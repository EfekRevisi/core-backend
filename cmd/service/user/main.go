package main

import (
	"core-backend/internal/service/userservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRoutes(router *gin.Engine) {
	router.GET("/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hais",
		})
	})
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(cors())
	controller := userservice.NewService()

	router.GET("/user", controller.GetAllUser)
	router.POST("/user", controller.InsertData)
	router.GET("/user/:id", controller.GetUserByID)
	router.PATCH("/user/:id", controller.UpdateUserByID)
	router.DELETE("/user/:id", controller.DeleteByID)

	router.Run()
}
