package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// the pq driver is used as an addon for the
	// basic sql library for golang
	_ "github.com/lib/pq"

	"github.com/fbFideles/fin-tracker/config"
)

var conf = config.GetConfig()

func newConnection() (database *sql.DB, err error) {
	connString := fmt.Sprintf("user=%v dbname=%v sslmode=%v",
		conf.Database.User,
		conf.Database.Name,
		conf.Database.SSL,
	)

	database, err = sql.Open("postgres", connString)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// NewTransaction defines a function with returns a new transaction
func NewTransaction(ctx context.Context) (transaction *sql.Tx, err error) {
	db, err := newConnection()
	if err != nil {
		log.Println(err)
		return
	}

	transaction, err = db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
