package service

import (
	"echo-project/model"
	"github.com/labstack/echo/v4"
)

type TestService interface {
	GetTests(c echo.Context) ([]*model.Test, error)
}
