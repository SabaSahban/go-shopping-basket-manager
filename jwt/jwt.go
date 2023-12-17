package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(userID int64, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token.Claims = claims
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		logrus.Errorf("failed to generate token: %s", err.Error())
		return "", err
	}

	return t, nil
}

func ValidateToken(tokenString string) (int64, error) {
	tkn, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if claims, ok := tkn.Claims.(*Claims); ok && tkn.Valid {
		logrus.Infof("validate token: User ID: %v, Username: %v", claims.UserID, claims.Username)
		return claims.UserID, nil
	}

	logrus.Errorf("error parsing token: %s", err.Error())
	return 0, errors.New("invalid token")
}
