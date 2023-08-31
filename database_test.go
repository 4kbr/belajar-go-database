package belajar_go_database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestHeloo(t *testing.T) {
	fmt.Println("Oke aja")
}

func TestOpenConnection(t *testing.T) {
	fmt.Println("Tezt")
	db, err := sql.Open("mysql", "root@tcp(localhost:8889)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected")
	defer db.Close()
	// menggunakan db
}
