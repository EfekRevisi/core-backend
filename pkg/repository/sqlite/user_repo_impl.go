package sqlite

import (
	"core-backend/pkg/models"
	"core-backend/pkg/repository"

	"github.com/jinzhu/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u *userRepositoryImpl) GetUserByID(ID string) (*models.User, error) {
	user := &models.User{}
	if error := u.db.First(&user, ID).Error; error != nil {
		return nil, error
	}

	return user, nil
}

func (u *userRepositoryImpl) GetAllUser() ([]models.User, error) {
	var users []models.User
	if error := u.db.Find(&users).Error; error != nil {
		return nil, error
	}

	return users, nil
}

func (u *userRepositoryImpl) UpdateUserByID(ID string, payload models.User) (*models.User, error) {
	user := &models.User{}

	if error := u.db.First(&user, ID).Error; error != nil {
		return nil, error
	}

	user.FirstName = payload.FirstName
	user.LastName = payload.LastName

	if error := u.db.Save(&user).Error; error != nil {
		return nil, error
	}

	return user, nil
}

func (u *userRepositoryImpl) DeleteByID(ID string) (bool, error) {
	user := &models.User{}

	if error := u.db.First(&user, ID).Error; error != nil {
		return false, error
	}

	if error := u.db.Delete(&user).Error; error != nil {
		return false, error
	}

	return true, nil
}
