package main

import (
	"ecommerce/config"
	productData "ecommerce/features/product/data"
	productHandler "ecommerce/features/product/handler"
	productService "ecommerce/features/product/services"
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

	productDt := productData.New(db)
	productSrv := productService.New(productDt)
	productHdl := productHandler.New(productSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	user := e.Group("/users")

	user.GET("", userHdl.GetData(), middleware.JWT([]byte(config.JWTKey)))
	user.GET("/:id", userHdl.Profile())
	user.PUT("", userHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	user.DELETE("", userHdl.Deactivate(), middleware.JWT([]byte(config.JWTKey)))

	e.GET("/products/:id", productHdl.GetById())
	e.GET("/products", productHdl.GetAll())
	e.POST("/products", productHdl.Add(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/products/:id", productHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/products/:id", productHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
