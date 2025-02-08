package repositories

import (
	"context"
	"errors"
	"fmt"
	errConstant "user-service/constants/error"
	"user-service/domain/dto"
	"user-service/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Register(context.Context, *dto.RegisterRequest) (*models.Users, error)
	// Login(context.Context, *dto.LoginRequest) (*models.Users, error)
	Update(context.Context, *dto.UpdateRequest, string) (*models.Users, error)
	// Delete(context.Context, string) error
	FindByUsername(context.Context, string) (*models.Users, error)
	FindByEmail(context.Context, string) (*models.Users, error)
	FindByUUID(context.Context, string) (*models.Users, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Register(ctx context.Context, req *dto.RegisterRequest) (*models.Users, error) {
	user := models.Users{
		UUID:        uuid.New(),
		RegNumber:   req.RegNumber,
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Photo:       req.Photo,
		RoleID:      req.RoleID,
		LibraryID:   req.LibraryID,
	}

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, req *dto.UpdateRequest, uuid string) (*models.Users, error) {
	user := models.Users{
		RegNumber:   req.RegNumber,
		Name:        req.Name,
		Username:    req.Username,
		Password:    *req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Photo:       req.Photo,
	}

	err := r.db.WithContext(ctx).
		Where("uuid = ?", uuid).
		Updates(user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", errConstant.ErrSQLError)
	}
	return &user, nil
}

// func (r *UserRepository) Delete(ctx context.Context, uuid string) error {
// 	userUUID := uuid
// 	return r.db.WithContext(ctx).Delete(&models.Users{}, "uuid = ?", userUUID).Error
// }

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.Users, error) {
	var user models.Users
	err := r.db.WithContext(ctx).
		Preload("Role").
		Preload("Library").
		Where("Username = ?", username).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstant.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.Users, error) {
	var user models.Users
	err := r.db.WithContext(ctx).
		Preload("Role").
		Preload("Library").
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstant.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) FindByUUID(ctx context.Context, uuid string) (*models.Users, error) {
	var user models.Users
	err := r.db.WithContext(ctx).
		Preload("Role").
		Preload("Library").
		Where("uuid = ?", uuid).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstant.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &user, nil
}
