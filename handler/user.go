package handler

import (
	"basketManager/jwt"
	"basketManager/model"
	"basketManager/request"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepo model.UserRepo
}

func (h *UserHandler) SignupHandler(c echo.Context) error {
	req := request.CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		logrus.Errorf("failed to bind signup request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := h.UserRepo.CreateUser(&user)
	if err != nil {
		logrus.Errorf("failed to create user: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, "User with username: "+
		user.Username+" created successfully")
}

func (h *UserHandler) LoginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.UserRepo.GetUserByUsername(username)
	if err != nil {
		logrus.Errorf("failed to get user by username: %s", err.Error())
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logrus.Errorf("password comparission error: %s", err.Error())
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}

	token, err := jwt.GenerateToken(user.ID, username)
	if err != nil {
		logrus.Errorf("failed to generate JWT token: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
