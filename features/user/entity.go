package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	Avatar   string
	Name     string
	Email    string
	Address  string
	Password string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Deactivate() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) error
	Login(email, password string) (string, Core, error)
	Profile(token interface{}) (interface{}, error)
	Update(formHeader multipart.FileHeader, token interface{}, updatedProfile Core) (Core, error)
	Deactivate(token interface{}) error
}

type UserData interface {
	Register(newUser Core) error
	Login(email string) (Core, error)
	Profile(id uint) (interface{}, error)
	Update(id uint, updatedProfile Core) (Core, error)
	Deactivate(id uint) error
}