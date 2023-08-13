package usecase

import (
	"gorestapi/model"
	"gorestapi/repository"
	"gorestapi/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user *model.User) (*model.UserResponse, error)
	Login(user *model.User) (string, error)
}

type UserUsecase struct {
	ur repository.IUserReposiroty
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserReposiroty, uv validator.IUserValidator) IUserUsecase {
	return &UserUsecase{
		ur: ur,
		uv: uv,
	}
}

func (uu *UserUsecase) SignUp(user *model.User) (*model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Email:    user.Email,
		Password: string(hash),
	}

	if err = uu.ur.Create(newUser); err != nil {
		return nil, err
	}

	userRes := &model.UserResponse{
		Id:    newUser.Id,
		Email: newUser.Email,
	}

	return userRes, nil
}

func (uu *UserUsecase) Login(user *model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	storedUser := &model.User{}
	if err := uu.ur.FindByEmail(storedUser, user.Email); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.Id,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
