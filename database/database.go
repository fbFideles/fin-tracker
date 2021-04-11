package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	// the pq driver is used as an addon for the
	// basic sql library for golang
	_ "github.com/lib/pq"

	"github.com/fbFideles/fin-tracker/config"
)

var (
	conf = config.GetConfig()
	db   *sql.DB
)

// NewConnection initializes the database conection
func NewConnection() (err error) {
	connString := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=%v",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Name,
		conf.Database.SSL,
	)

	db, err = sql.Open("postgres", connString)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// CloseConnection closes the active database section
func CloseConnection() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// NewTransaction defines a function with returns a new transaction
func NewTransaction(ctx context.Context) (transaction *sql.Tx, err error) {
	if db == nil {
		return nil, errors.New("Database connection is not initialized please start")
	}

	transaction, err = db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
