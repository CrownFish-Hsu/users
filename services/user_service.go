package services

import (
	"blog/domain/users"
	"blog/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to bcrypt password.")
	}

	user.Password = string(pwd[:])
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{Email: user.Email}

	if err := result.GetByEmail(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, errors.NewBadRequestError("failed to decrypt password.")
	}

	return result, nil
}

func GetUserByID(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}

	if err := result.GetById(); err != nil {
		return nil, err
	}

	return result, nil
}