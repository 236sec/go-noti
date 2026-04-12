package models

type User struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Username    string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Role        string `gorm:"not null"`
	DateOfBirth string `gorm:"type:date"`
}
