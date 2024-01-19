package repository

import (
	"echo-project/model"
	"gorm.io/gorm"
)

type TestRepositoryImpl struct {
	db *gorm.DB
}

func (TestRepositoryImpl) NewTestRepositoryImpl(db *gorm.DB) *TestRepositoryImpl {
	return &TestRepositoryImpl{db}
}

func (testRepositoryImpl *TestRepositoryImpl) FindAll() ([]*model.Test, error) {
	var tests []*model.Test

	err := testRepositoryImpl.db.Find(&tests).Error
	if err != nil {
		return nil, err
	}

	return tests, nil
}
