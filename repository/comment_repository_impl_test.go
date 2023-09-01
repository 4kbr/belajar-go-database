package repository

import (
	belajar_go_database "belajar-go-database"
	"belajar-go-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_go_database.GetConnection())
	ctx := context.Background()

	comment := entity.Comment{
		Email:   "testrepo@test.com",
		Comment: "bismillah ges",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println("result", result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_go_database.GetConnection())
	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 5) // ini test bisa
	// comment, err := commentRepository.FindById(ctx, 102) // ini test gk bisa

	if err != nil {
		panic(err)
	}

	fmt.Println("result", comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_go_database.GetConnection())
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
	// fmt.Println("result", comment)
}
