package db

import (
	// "database/sql"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var PGURL = "postgresql://jxyz:1234@localhost/jr"

func GetPG() *sqlx.DB {
	db, err := sqlx.Connect("pgx", PGURL)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
