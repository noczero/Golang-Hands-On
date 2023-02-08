package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/noczero/Golang-Hands-On/Database"
	"github.com/noczero/Golang-Hands-On/Database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	// invoke repository adapter
	commentRepository := NewCommentRepository(Database.GetConnection())

	// prepare function statement
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "code@test.com",
		Comment: "Lorem Ipsum Test",
	}

	// invoke insert function
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(Database.GetConnection())
	result, err := commentRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(Database.GetConnection())
	result, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
