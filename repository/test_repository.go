package repository

import (
	"echo-project/model"
)

type TestRepository interface {
	FindAll() ([]*model.Test, error)
}
