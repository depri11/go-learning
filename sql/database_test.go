package sql

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestConnectDB(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=123 dbname=learn sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
