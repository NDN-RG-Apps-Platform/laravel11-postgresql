package services

import (
	"context"
	"strings"
	"time"
	"user-service/config"
	"user-service/constants"
	errConstant "user-service/constants/error"
	"user-service/domain/dto"
	"user-service/domain/models"
	"user-service/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repositories.IRespositoryRegistry
}

type IUserService interface {
	Login(context.Context, *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(context.Context, *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Update(context.Context, *dto.UpdateRequest, string) (*dto.UserResponse, error)
	GetUserLogin(context.Context) (*dto.UserResponse, error)
	GetUserByUUID(context.Context, string) (*dto.UserResponse, error)
}

type Claims struct {
	User *dto.UserResponse
	jwt.RegisteredClaims
}

func NewUserService(repository repositories.IRespositoryRegistry) IUserService {
	return &UserService{repository: repository}
}

func (u *UserService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// cari user berdasarkan username
	user, err := u.repository.GetUser().FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	// set expiraion time
	expirationTime := time.Now().Add(time.Duration(config.Config.JwtExpirationTime) * time.Minute).Unix()

	// buat data user response
	data := &dto.UserResponse{
		UUID:        user.UUID,
		RegNumber:   user.RegNumber,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Photo:       user.Photo,
		Role:        strings.ToLower(user.Role.Code),
		Library:     strings.ToLower(user.Library.Code),
	}

	// klaim jwt token
	claims := &Claims{
		User: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JwtSecretKey))
	if err != nil {
		return nil, err
	}

	// response login
	response := &dto.LoginResponse{
		Token: tokenString,
		User:  *data,
	}

	return response, nil
}

func (u *UserService) isUsernameExist(ctx context.Context, username string) bool {
	user, err := u.repository.GetUser().FindByUsername(ctx, username)
	if err != nil {
		return false
	}
	if user != nil {
		return true
	}
	return false
}

func (u *UserService) isEmailExist(ctx context.Context, email string) bool {
	user, err := u.repository.GetUser().FindByEmail(ctx, email)
	if err != nil {
		return false
	}
	if user != nil {
		return true
	}
	return false
}

func (u *UserService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if u.isUsernameExist(ctx, req.Username) {
		return nil, errConstant.ErrUsernameExist
	}

	if u.isEmailExist(ctx, req.Email) {
		return nil, errConstant.ErrEmailExist
	}

	if req.Password != req.ConfirmPassword {
		return nil, errConstant.ErrPasswordDoesNotMatch
	}

	user, err := u.repository.GetUser().Register(ctx, &dto.RegisterRequest{
		RegNumber:   req.RegNumber,
		Name:        req.Name,
		Username:    req.Username,
		Password:    string(hashedPassword),
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Photo:       req.Photo,
		RoleID:      req.RoleID,
		LibraryID:   req.LibraryID,
	})

	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		User: dto.UserResponse{
			UUID:        user.UUID,
			RegNumber:   user.RegNumber,
			Name:        user.Name,
			Username:    user.Username,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Photo:       user.Photo,
		},
	}

	return response, nil
}

func (u *UserService) Update(ctx context.Context, req *dto.UpdateRequest, uuid string) (*dto.UserResponse, error) {
	var (
		password                  string
		checkUsername, checkEmail *models.Users
		hashedPassword            []byte
		user, userResult          *models.Users
		err                       error
		data                      dto.UserResponse
	)

	user, err = u.repository.GetUser().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	isUsernameExits := u.isUsernameExist(ctx, req.Username)
	if isUsernameExits && user.Username != req.Username {
		checkUsername, err = u.repository.GetUser().FindByUsername(ctx, req.Username)
		if err != nil {
			return nil, err
		}

		if checkUsername != nil {
			return nil, errConstant.ErrUsernameExist
		}
	}

	isEmailExits := u.isEmailExist(ctx, req.Email)
	if isEmailExits && user.Email != req.Email {
		checkEmail, err = u.repository.GetUser().FindByEmail(ctx, req.Email)
		if err != nil {
			return nil, err
		}

		if checkEmail != nil {
			return nil, errConstant.ErrEmailExist
		}
	}

	if req.Password != nil {
		if *req.Password != *req.ConfirmPassword {
			return nil, errConstant.ErrPasswordDoesNotMatch
		}

		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		password = string(hashedPassword)
	}

	userResult, err = u.repository.GetUser().Update(ctx, &dto.UpdateRequest{
		RegNumber:   req.RegNumber,
		Name:        req.Name,
		Username:    req.Username,
		Password:    &password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Photo:       req.Photo,
	}, uuid)

	if err != nil {
		return nil, err
	}

	data = dto.UserResponse{
		UUID:        userResult.UUID,
		RegNumber:   userResult.RegNumber,
		Name:        userResult.Name,
		Username:    userResult.Username,
		Email:       userResult.Email,
		PhoneNumber: userResult.PhoneNumber,
		Photo:       userResult.Photo,
	}

	return &data, nil
}

func (u *UserService) GetUserLogin(ctx context.Context) (*dto.UserResponse, error) {
	var (
		userLogin = ctx.Value(constants.UserLogin).(*dto.UserResponse)
		data      dto.UserResponse
	)

	data = dto.UserResponse{
		UUID:        userLogin.UUID,
		RegNumber:   userLogin.RegNumber,
		Name:        userLogin.Name,
		Username:    userLogin.Username,
		Email:       userLogin.Email,
		PhoneNumber: userLogin.PhoneNumber,
		Photo:       userLogin.Photo,
		Role:        userLogin.Role,
		Library:     userLogin.Library,
	}

	return &data, nil
}

func (u *UserService) GetUserByUUID(ctx context.Context, uuid string) (*dto.UserResponse, error) {
	user, err := u.repository.GetUser().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	data := dto.UserResponse{
		UUID:        user.UUID,
		RegNumber:   user.RegNumber,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Photo:       user.Photo,
	}

	return &data, nil
}
