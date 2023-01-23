package service

import (
	"github.com/noczero/Golang-Hands-On/UnitTest/entity"
	"github.com/noczero/Golang-Hands-On/UnitTest/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

/*
Implements mock test in category service
*/

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository} // using mock as service

func TestCategoryService_GetNotFound(t *testing.T) {

	// create mock data using arguments, and return nil
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// get mock data by arguments
	category, err := categoryService.Get("1")
	// test nil
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategorySerivce_Found(t *testing.T) {
	expectedCategory := entity.Category{
		Id:   "01",
		Name: "Satrya",
	}

	// create mock data using arguments, return expected result
	categoryRepository.Mock.On("FindById", "2").Return(expectedCategory)

	// get mock data by arguments using the service
	result, err := categoryService.Get("2")

	// test result
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, expectedCategory.Id, result.Id)
	assert.Equal(t, expectedCategory.Name, result.Name)
}
