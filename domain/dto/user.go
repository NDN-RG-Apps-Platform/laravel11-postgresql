package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" validates:"required"`
	Password string `json:"password" validates:"required"`
}

type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	RegNumber   string    `json:"regNumber"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Role        string    `json:"role,omitempty"`
	Library     string    `json:"library,omitempty"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type RegisterRequest struct {
	RegNumber       string `json:"regNumber" validates:"required"`
	Name            string `json:"name" validates:"required"`
	Username        string `json:"username" validates:"required"`
	Password        string `json:"password" validates:"required"`
	ConfirmPassword string `json:"confirmPassword" validates:"required"`
	Email           string `json:"email" validates:"required,email"`
	PhoneNumber     string `json:"phoneNumber" validates:"required"`
	RoleID          uint
	LibraryID       uint
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}

type UpdateRequest struct {
	RegNumber       string `json:"regNumber" validates:"required"`
	Name            string `json:"name" validates:"required"`
	Username        string `json:"username" validates:"required"`
	Password        string `json:"password" validates:"required"`
	ConfirmPassword string `json:"confirmPassword" validates:"required"`
	Email           string `json:"email" validates:"required,email"`
	PhoneNumber     string `json:"phoneNumber" validates:"required"`
	RoleID          uint
	LibraryID       uint
}
