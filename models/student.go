package models

import "github.com/google/uuid"

type Student struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	StudentCode string `json:"studentCode" gorm:"not null;unique;type:varchar(20)"`
	Prefix string `json:"prefix" gorm:"type:varchar"`
	FirstName string `json:"firstName" gorm:"not null;type:varchar(30)"`
	LastName string `json:"lastName" gorm:"not null;type:varchar(30)"`
	Email string `json:"email" gorm:"type:varchar(50)"`
	Phone string `json:"phone" gorm:"type:varchar(20)"`
}