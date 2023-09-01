package belajar_go_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
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

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	// script := "select id, name from customer" bisa juga
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err =
			rows.Scan(
				&id,
				&name,
				&email,
				&balance,
				&rating,
				&birthDate,
				&married,
				&createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Print("Id:", id, " ")
		fmt.Print("Name:", name, " ")
		if email.Valid {
			fmt.Print("email:", email.String, " ")
		}
		fmt.Print("balance:", balance, " ")
		fmt.Print("rating:", rating, " ")
		if birthDate.Valid {
			fmt.Print("birthDate:", birthDate.Time, " ")
		}
		fmt.Print("married:", married, " ")
		fmt.Print("createdAt:", createdAt, " ")
		fmt.Println("")
		fmt.Println("-")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin1'; #"
	password := "contohpasswordkautsss"

	//bahaya sql injection
	script := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin1"
	password := "contohpasswordkaut"

	//safe dari sql injection
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin2; DROP TABLE user; #"
	password := "pw2"

	script := "INSERT INTO user(username,password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succcess Insert new Cutomer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	email := "aku@gmail.com"
	comment := "Mantap bos~~~"

	script := "INSERT INTO comments(email,comment) VALUES(?,?)"

	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succcess Insert new comment with id:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	script := "INSERT INTO comments(email,comment) VALUES(?,?)"

	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 10; i < 20; i++ {
		email := "ke" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Commetnt ke", id)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	//TODO: transaction
	script := "INSERT INTO comments(email,comment) VALUES(?,?)"

	for i := 0; i < 5; i++ {
		email := "ke" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ke", id)
	}

	// err = tx.Commit()
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}
