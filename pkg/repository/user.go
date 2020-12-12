package repository

import "core-backend/pkg/models"

// UserRepository - for all access user repo
type UserRepository interface {
	GetUserByID(ID string) (*models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateUserByID(ID string, payload models.User) (*models.User, error)
	DeleteByID(ID string) (bool, error)
}
