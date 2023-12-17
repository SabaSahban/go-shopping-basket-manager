package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64      `json:"id" gorm:"primary_key"` // Auto generated ID
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"-"`
}

type UserRepo interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}

type SQLUserRepo struct {
	DB *gorm.DB
}

func NewSQLUserRepo(db *gorm.DB) *SQLUserRepo {
	return &SQLUserRepo{DB: db}
}

func (r *SQLUserRepo) CreateUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return r.DB.Create(user).Error
}

func (r *SQLUserRepo) GetUserByUsername(username string) (*User, error) {
	user := &User{}

	err := r.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
