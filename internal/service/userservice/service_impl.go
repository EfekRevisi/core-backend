package userservice

import (
	"core-backend/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceImpl struct {
	registry ServiceRegistry
}

// NewService returns service instance.
func NewService() UserService {
	return &serviceImpl{
		registry: NewServiceRegistry(),
	}
}

func (service *serviceImpl) GetUserByID(c *gin.Context) {
	data, err := service.registry.GetUserReadRepo().GetUserByID(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}

func (service *serviceImpl) GetAllUser(c *gin.Context) {
	data, err := service.registry.GetUserReadRepo().GetAllUser()

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}

func (service *serviceImpl) InsertData(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	user := models.User{
		ID:        input.ID,
		FirstName: input.Firstname,
		LastName:  input.Lastname,
	}

	data, err := service.registry.GetUserReadRepo().Insert(user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (service *serviceImpl) UpdateUserByID(c *gin.Context) {
	data, err := service.registry.GetUserReadRepo().GetUserByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := service.registry.GetUserReadRepo().UpdateUserByID(
		c.Param("id"),
		data,
		input,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (service *serviceImpl) DeleteByID(c *gin.Context) {
	data, err := service.registry.GetUserReadRepo().DeleteByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
