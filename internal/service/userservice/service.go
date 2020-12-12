package userservice

import "github.com/gin-gonic/gin"

// UserService interface for user service
type UserService interface {
	GetUserByID(c *gin.Context)

	GetAllUser(c *gin.Context)

	InsertData(c *gin.Context)

	UpdateUserByID(c *gin.Context)

	DeleteByID(c *gin.Context)
}
