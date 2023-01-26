package main

import (
	"ecommerce/config"
	cartData "ecommerce/features/cart/data"
	cartHandler "ecommerce/features/cart/handler"
	cartService "ecommerce/features/cart/services"
	productData "ecommerce/features/product/data"
	productHandler "ecommerce/features/product/handler"
	productService "ecommerce/features/product/services"
	trxD "ecommerce/features/transaction/data"
	trxH "ecommerce/features/transaction/handler"
	trxS "ecommerce/features/transaction/services"
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

	cartDt := cartData.New(db)
	cartSrv := cartService.New(cartDt)
	cartHdl := cartHandler.New(cartSrv)

	trxData := trxD.New(db)
	trxSrv := trxS.New(trxData)
	trxHdl := trxH.New(trxSrv)

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

	e.GET("/products/:id", productHdl.GetById())
	e.GET("/products", productHdl.GetAll())
	e.POST("/products", productHdl.Add(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/products/:id", productHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/products/:id", productHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))

	e.POST("/carts/:idProduct", cartHdl.Add(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/carts/:idCart", cartHdl.GetByIdC(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/carts", cartHdl.GetByIdU(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/carts/:idCart", cartHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/carts/:idCart", cartHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))

	e.POST("/transactions/:id", trxHdl.Add(), middleware.JWT([]byte(config.JWTKey)))
	// e.GET("/products/:idUser", cartHdl.GetByIdU(), middleware.JWT([]byte(config.JWTKey)))
	// e.GET("/products", cartHdl.GetAll())
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
