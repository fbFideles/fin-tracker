package database

import (
	"database/sql"
	"fmt"
	"log"

	// the pq driver is used as an addon for the
	// basic sql library for golang
	_ "github.com/lib/pq"

	"github.com/fbFideles/fin-tracker/config"
)

var conf = config.GetConfig()

// NewConnection opens a new connection with the database
func NewConnection() (database *sql.DB, err error) {
	connString := fmt.Sprintf("user=%v dbname=%v sslmode=%v",
		conf.Database.User,
		conf.Database.Name,
		conf.Database.SSL,
	)

	database, err = sql.Open("postgres", connString)
	if err != nil {
		log.Println(err)
	}
	return
}