package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("postgres", "user=test password=password dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	exitCode := m.Run()

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

	os.Exit(exitCode)
}

func TestItemRepository_Save(t *testing.T) {

}
