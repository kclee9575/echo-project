package service

import (
	"echo-project/model"
	"echo-project/repository"
	"github.com/labstack/echo/v4"
)

type TestServiceImpl struct {
	repository.TestRepository
}

func (TestServiceImpl) NewTestServiceImpl(testRepository repository.TestRepository) *TestServiceImpl {
	return &TestServiceImpl{testRepository}
}

func (TestServiceImpl *TestServiceImpl) GetTests(c echo.Context) ([]*model.Test, error) {
	return TestServiceImpl.TestRepository.FindAll()
}
