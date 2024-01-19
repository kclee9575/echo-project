package route

import (
	"echo-project/config"
	"echo-project/handler"
	"echo-project/repository"
	"echo-project/route/middlewares"
	"echo-project/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) (*echo.Echo, error) {
	e.Use(middlewares.LoggingMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// TODO
	// 아직 미완성..
	// route 어떻게 할지 고민중입니다.
	db := config.GetDataBase()
	testRepository := repository.TestRepositoryImpl{}.NewTestRepositoryImpl(db)
	testService := service.TestServiceImpl{}.NewTestServiceImpl(testRepository)
	testHandler := handler.TestHandler{}.NewTestHandler(testService)

	testHandler.Init(e.Group("/api/v1"))

	return e, nil
}
