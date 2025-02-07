package models

import "time"

type Libraries struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Code        string `gorm:"type:varchar(15);not null"`
	LibraryName string `gorm:"type:varchar(30);not null"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
