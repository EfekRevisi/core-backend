package models

// User user model
type User struct {
	ID        int32  `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	FirstName string `gorm:"not null" form:"firstname" json:"firstname"`
	LastName  string `gorm:"not null" form:"lastname" json:"lastname"`
}
