package server

import (
	"basketManager/config"
	"basketManager/db"
	"basketManager/handler"
	"basketManager/jwt"
	"basketManager/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "start a new server for basket manager app",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	})
}

func main(cfg config.Config) {
	database := db.WithRetry(db.Create, cfg.Postgres)
	defer func() {
		if err := database.Close(); err != nil {
			logrus.Error(err.Error())
		}
	}()

	basketRepo := model.NewSQLBasketRepo(database)
	userRepo := model.NewSQLUserRepo(database)

	basketHandler := handler.BasketHandler{
		BasketRepo: basketRepo,
	}

	userHandler := handler.UserHandler{
		UserRepo: userRepo,
	}

	e := echo.New()

	e.POST("/signup", userHandler.SignupHandler)
	e.POST("/login", userHandler.LoginHandler)

	api := e.Group("/api")

	api.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:    middleware.DefaultSkipper,
		KeyLookup:  "header:" + echo.HeaderAuthorization,
		AuthScheme: "Bearer",
		Validator: func(key string, context echo.Context) (b bool, e error) {
			id, err := jwt.ValidateToken(key)
			if err != nil {
				return false, err
			}

			context.Set("user_id", id)

			return true, nil
		},
	}))

	api.GET("/basket", basketHandler.GetAll)
	api.POST("/basket", basketHandler.Create)
	api.PATCH("/basket/:id", basketHandler.Update)
	api.GET("/basket/:id", basketHandler.GetByBasketID)
	api.DELETE("/basket/:id", basketHandler.Delete)

	port := cfg.Server.Port
	serverAddress := fmt.Sprintf(":%d", port)
	if err := e.Start(serverAddress); !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("server failed with error: %v", err)
	}
}
