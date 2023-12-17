package model

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Basket struct {
	ID        int64           `json:"id" gorm:"primary_key"` // Auto generated ID
	Data      json.RawMessage `json:"data"`
	State     string          `json:"state"`
	UserID    int64           `json:"user_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt *time.Time      `json:"-"`
}

type BasketRepo interface {
	CreateBasket(basket *Basket) error
	GetBasketByID(userID, id int64) (*Basket, error)
	UpdateBasket(userID int64, basket *Basket) error
	DeleteBasket(userID, id int64) error
	GetAllBaskets(userID int64) ([]*Basket, error)
}

type SQLBasketRepo struct {
	DB *gorm.DB
}

func NewSQLBasketRepo(db *gorm.DB) *SQLBasketRepo {
	return &SQLBasketRepo{DB: db}
}

func (r *SQLBasketRepo) CreateBasket(basket *Basket) error {
	return r.DB.Create(basket).Error
}

func (r *SQLBasketRepo) GetBasketByID(userID, id int64) (*Basket, error) {
	basket := &Basket{}
	err := r.DB.Where("user_id = ? AND id = ?", userID, id).First(basket).Error
	if err != nil {
		return nil, err
	}

	return basket, nil
}

func (r *SQLBasketRepo) UpdateBasket(userID int64, basket *Basket) error {
	existingBasket, err := r.GetBasketByID(userID, basket.ID)
	if err != nil {
		return err
	}

	if existingBasket.State == "COMPLETED" {
		return errors.New("cannot update a basket with state COMPLETED")
	}

	existingBasket.Data = basket.Data
	existingBasket.State = basket.State

	if err := r.DB.Save(existingBasket).Error; err != nil {
		return err
	}

	return nil
}

func (r *SQLBasketRepo) DeleteBasket(userID, id int64) error {
	return r.DB.Where("user_id = ? AND id = ?", userID, id).Delete(&Basket{}).Error
}

func (r *SQLBasketRepo) GetAllBaskets(userID int64) ([]*Basket, error) {
	var baskets []*Basket
	if err := r.DB.Where("user_id = ?", userID).Find(&baskets).Error; err != nil {
		return nil, err
	}

	return baskets, nil
}
