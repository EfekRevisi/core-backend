package userservice

import (
	"core-backend/pkg/repository"
	"core-backend/pkg/repository/sqlite"
)

// ServiceRegistry interface.
type ServiceRegistry interface {
	GetUserReadRepo() repository.UserRepository
}

type serviceRegistryImpl struct {
	up *repository.UserRepository
}

//NewServiceRegistry returns ServiceRegistry instance.
func NewServiceRegistry() ServiceRegistry {
	return &serviceRegistryImpl{}
}

func createUserRepo() interface{} {
	return sqlite.InitUserRepository(sqlite.ConnectDatabase())
}

func (s *serviceRegistryImpl) GetUserReadRepo() repository.UserRepository {
	if s.up != nil {
		return *s.up
	}

	return sqlite.InitUserRepository(sqlite.ConnectDatabase())
}
