package repository

import (
	"github.com/noczero/Golang-Hands-On/UnitTest/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id) // define mock based on the function parameter

	if arguments.Get(0) == nil {
		return nil
	} else {
		// conversion interface to origin object
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
