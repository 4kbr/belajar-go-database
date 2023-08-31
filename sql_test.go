package belajar_go_database

import (
	"context"
	"fmt"
	"testing"
)

// test input ke mysql
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('kita','Kita')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succcess Insert new Cutomer")
}

// test get data dari mysql
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	// script := "select id, name from customer" bisa juga
	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Print("Id:", id)
		fmt.Print(" -  ")
		fmt.Print("Name:", name)
		fmt.Println("")

	}
}
