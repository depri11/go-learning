package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := SetupDB()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(name) VALUES ('Devri')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}

func TestQuerySql(t *testing.T) {
	db := SetupDB()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := SetupDB()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id,name,email,balance,rating,birth_date,married,created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email sql.NullString
		var balance int
		var rating float32
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date: ", birthDate.Time)
		}
		fmt.Println("Married: ", married)
		fmt.Println("CreatedAt: ", createdAt)
	}
}

func TestSqlInjections(t *testing.T) {
	db := SetupDB()
	defer db.Close()

	username := "user'; //"
	password := "123"

	ctx := context.Background()
	script := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
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
		fmt.Println("Success Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionsSafe(t *testing.T) {
	db := SetupDB()
	defer db.Close()

	username := "users"
	password := "123"

	ctx := context.Background()
	script := "SELECT username FROM users WHERE username = $1 AND password = $2 LIMIT 1"
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
		fmt.Println("Success Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestPrepareStatement(t *testing.T) {
	db := SetupDB()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email,comment) VALUES ($1,$2)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "dev" + strconv.Itoa(i) + "email"
		comment := "komentar ke: " + strconv.Itoa(i) + "comment"

		_, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
	}
}

func TestTransaction(t *testing.T) {
	db := SetupDB()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email,comment) VALUES ($1,$2)"

	for i := 0; i < 10; i++ {
		email := "dev" + strconv.Itoa(i) + "email"
		comment := "komentar ke: " + strconv.Itoa(i) + "comment"

		_, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
