package repositories

import (
	repositories "user-service/repositories/user"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRespositoryRegistry interface {
	GetUser() repositories.IUserRepository
}

func NewRespositoryRegistry(db *gorm.DB) IRespositoryRegistry {
	return &Registry{db: db}
}

func (r *Registry) GetUser() repositories.IUserRepository {
	return repositories.NewUserRepository(r.db)
}
