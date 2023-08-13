package controller

import (
	"gorestapi/model"
	"gorestapi/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CSRFToken(c echo.Context) error
}

type UserController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &UserController{
		uu: uu,
	}
}

func (uc *UserController) SignUp(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (uc *UserController) Login(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	jwt, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookei := &http.Cookie{
		Name:     "token",
		Value:    jwt,
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	c.SetCookie(cookei)

	return c.NoContent(http.StatusOK)
}

func (uc *UserController) Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Expires:  time.Now(),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (uc *UserController) CSRFToken(c echo.Context) error {
	token := c.Get("csrf").(string)

	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
