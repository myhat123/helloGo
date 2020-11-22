package main

import (
	// "database/sql"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"hello_17/common"
	"hello_17/tasks"
)

func main() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("pgx", "postgresql://jxyz:1234@localhost/jr")
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Queryx("SELECT * FROM brch_qry_dtl")

	records := make([]*common.DBrchQryDtl, 0)
	i := 0
	for rows.Next() {
		d := new(common.DBrchQryDtl)

		err := rows.StructScan(&d)
		if err != nil {
			log.Fatalln(err)
		}

		records = append(records, d)

		i = i + 1
	}

	tasks.Start(db, records)
}
