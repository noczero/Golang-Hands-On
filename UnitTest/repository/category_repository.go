package repository

import "github.com/noczero/Golang-Hands-On/UnitTest/entity"

/*
THis layer represent interface for manipulate data through database connect to entity layer
*/

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
