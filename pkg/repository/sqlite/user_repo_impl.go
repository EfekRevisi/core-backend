package sqlite

import (
	"core-backend/pkg/models"
	"core-backend/pkg/repository"

	"github.com/jinzhu/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

// InitUserRepository - for initialize user repository
func InitUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u *userRepositoryImpl) GetUserByID(ID string) (*models.User, error) {
	user := &models.User{}
	if err := u.db.First(&user, ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepositoryImpl) GetAllUser() ([]models.User, error) {
	var users []models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepositoryImpl) Insert(payload models.User) (*models.User, error) {
	user := &models.User{
		ID:        payload.ID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepositoryImpl) UpdateUserByID(
	ID string,
	user *models.User,
	payload models.UpdateUserInput,
) (bool, error) {

	user.FirstName = payload.Firstname
	user.LastName = payload.Lastname

	if err := u.db.Save(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (u *userRepositoryImpl) DeleteByID(ID string) (bool, error) {
	user := &models.User{}

	if err := u.db.First(&user, ID).Error; err != nil {
		return false, err
	}

	if err := u.db.Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
