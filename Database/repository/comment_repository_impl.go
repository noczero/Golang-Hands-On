package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/noczero/Golang-Hands-On/Database/entity"
	"strconv"
)

// implementation for comment repository contract

type commentRepositoryImpl struct {
	DB *sql.DB
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	//implement Insert
	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment) // execute query

	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id) // update comment id from database
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	//implement FindById
	query := "SELECT id,email,comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)

	// make empty entity
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close() // if using Query context, dont forget to close.

	if rows.Next() {
		// convert to entity
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, err
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " was not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	//Implement FindAll
	query := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	// make slice of comment to append
	var comments []entity.Comment

	for rows.Next() {
		comment := entity.Comment{}                              // make empty comment
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment) // fill the value of empty comment from scan

		// append comment to slice of comments
		comments = append(comments, comment)
	}

	return comments, err
}

// entrypoint of adapter
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db} // generate automatically
}
