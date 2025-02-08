package services

import (
	"user-service/repositories"
	services "user-service/services/user"
)

type Registry struct {
	repository repositories.IRespositoryRegistry
}

type IserviceRegistry interface {
	GetUser() services.IUserService
}

func NewServiceRegistry(repository repositories.IRespositoryRegistry) IserviceRegistry {
	return &Registry{repository: repository}
}

func (r *Registry) GetUser() services.IUserService {
	return services.NewUserService(r.repository)
}
