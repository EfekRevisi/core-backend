package models

// User user model
type User struct {
	ID        int32  `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	FirstName string `gorm:"not null" form:"firstname" json:"firstname"`
	LastName  string `gorm:"not null" form:"lastname" json:"lastname"`
}

// CreateUserInput create user input json validator
type CreateUserInput struct {
	ID        int32  `json:"id" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

// UpdateUserInput update user input json validator
type UpdateUserInput struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
