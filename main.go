package main

import (
	"ecommerce/config"
	usrD "ecommerce/features/user/data"
	usrH "ecommerce/features/user/handler"
	usrS "ecommerce/features/user/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := usrD.New(db)
	userSrv := usrS.New(userData)
	userHdl := usrH.New(userSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	user := e.Group("/users")

	user.GET("", userHdl.Profile(), middleware.JWT([]byte(config.JWTKey)))
	user.PUT("", userHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	user.DELETE("", userHdl.Deactivate(), middleware.JWT([]byte(config.JWTKey)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
