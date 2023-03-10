package base

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {

	if err != nil {
		logFatal(err)
	}
}

func ConnectDB() *sql.DB {

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)
	log.Println(pgUrl)
	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)
	err = db.Ping()
	logFatal(err)
}
