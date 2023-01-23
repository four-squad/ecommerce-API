package helper

import (
	"github.com/labstack/echo/v4"
)

type MediaDto struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Data       *echo.Map `json:"data"`
}
