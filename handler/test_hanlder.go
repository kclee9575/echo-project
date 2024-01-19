package handler

import (
	"echo-project/service"
	"github.com/labstack/echo/v4"
)

type TestHandler struct {
	service.TestService
}

func (TestHandler) NewTestHandler(service service.TestService) *TestHandler {
	return &TestHandler{service}
}

func (testHandler *TestHandler) Init(e *echo.Group) {
	e.GET("/test", testHandler.GetTests)
}

// GetTests get all tests' list
// @Summary Get all tests
// @Description Get all test's info
// @Accept json
// @Produce json
// @Success 200 {object} ApiResult{result=model.Test}
// @Router /tests [get]
func (testHandler *TestHandler) GetTests(c echo.Context) error {
	tests, err := testHandler.TestService.GetTests(c)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, tests)
}
