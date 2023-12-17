package handler

import (
	"basketManager/request"
	"errors"
	"net/http"
	"strconv"

	"basketManager/model"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type BasketHandler struct {
	BasketRepo model.BasketRepo
}

func (h *BasketHandler) Create(c echo.Context) error {
	req := request.CreateBasketRequest{}

	if err := c.Bind(&req); err != nil {
		logrus.Errorf("create basket handler binding request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	userID, err := h.getUserID(c)
	if err != nil {
		logrus.Errorf("create basket handler get user ID: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	basket := model.Basket{
		Data:   req.Data,
		State:  req.State,
		UserID: userID,
	}

	if err := h.BasketRepo.CreateBasket(&basket); err != nil {
		logrus.Errorf("create basket handler failed: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create basket")
	}

	return c.JSON(http.StatusOK, basket)
}

func (h *BasketHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Errorf("update basket handler get basket ID: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid basket ID")
	}

	req := request.UpdateBasketRequest{}

	if err := c.Bind(&req); err != nil {
		logrus.Errorf("update basket handler binding request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	basket := model.Basket{
		ID:    int64(id),
		Data:  req.Data,
		State: req.State,
	}

	userID, err := h.getUserID(c)
	if err != nil {
		logrus.Errorf("update basket handler get user ID: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	err = h.BasketRepo.UpdateBasket(userID, &basket)
	if err != nil {
		logrus.Errorf("failed to update basket: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update basket")
	}

	return c.JSON(http.StatusOK, basket)
}

func (h *BasketHandler) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	userID, err := h.getUserID(c)
	if err != nil {
		logrus.Errorf("delete basket handler get user ID: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	if err := h.BasketRepo.DeleteBasket(userID, int64(id)); err != nil {
		logrus.Errorf("failed to delete basket: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete basket")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *BasketHandler) GetAll(c echo.Context) error {
	userID, err := h.getUserID(c)
	if err != nil {
		logrus.Errorf("get all basket handler get user ID: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	baskets, err := h.BasketRepo.GetAllBaskets(userID)
	if err != nil {
		logrus.Errorf("failed to get all baskets: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get baskets")
	}

	return c.JSON(http.StatusOK, baskets)
}

func (h *BasketHandler) GetByBasketID(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	userID, err := h.getUserID(c)
	if err != nil {
		logrus.Errorf("get by basket id handler get user ID: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	basket, err := h.BasketRepo.GetBasketByID(userID, int64(id))
	if err != nil {
		logrus.Errorf("failed to get basket by ID: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get basket")
	}

	return c.JSON(http.StatusOK, basket)
}

func (h *BasketHandler) getUserID(c echo.Context) (int64, error) {
	var userID int64

	uID := c.Get("user_id")

	if userIDInt, ok := uID.(int64); ok {
		userID = userIDInt
	} else {
		return 0, errors.New("failed to extract user ID")
	}

	return userID, nil
}
