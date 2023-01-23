package service

import (
	"errors"
	"github.com/noczero/Golang-Hands-On/UnitTest/entity"
	"github.com/noczero/Golang-Hands-On/UnitTest/repository"
)

/*
This later represent business logic layer connectiong to repository layer
*/

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
