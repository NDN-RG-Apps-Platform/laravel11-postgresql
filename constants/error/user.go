package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordInCorrect    = errors.New("password incorect")
	ErrUsernameExist        = errors.New("username already exits")
	ErrEmailExist           = errors.New("email already exits")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
}
