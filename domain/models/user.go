package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UUID        uuid.UUID `gorm:"type:uuid;not null"`
	RegNumber   string    `gorm:"type:varchar(15);not null;unique"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Username    string    `gorm:"type:varchar(50);not null;unique"`
	Password    string    `gorm:"type:varchar(100);not null"`
	Email       string    `gorm:"type:varchar(100);not null;unique"`
	PhoneNumber string    `gorm:"type:varchar(15);not null"`
	Photo       string    `gorm:"type:varchar(225);not null"`
	RoleID      uint      `gorm:"type:uint;not null"`
	LibraryID   uint      `gorm:"type:uint;not null"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Role        Roles     `gorm:"foreignKey:role_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Library     Libraries `gorm:"foreignKey:library_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
