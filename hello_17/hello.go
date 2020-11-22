package main

import (
	// "database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	// _ "github.com/lib/pq"
)

type DBrchQryDtl struct {
	Acc        string
	TranDate   time.Time `db:"tran_date"`
	Amt        string
	DrCrFlag   int    `db:"dr_cr_flag"`
	RptSum     string `db:"rpt_sum"`
	Timestamp1 string
}

func main() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("pgx", "postgresql://jxyz:1234@localhost/jr")
	if err != nil {
		log.Fatalln(err)
	}

	d := DBrchQryDtl{}
	rows, err := db.Queryx("SELECT * FROM brch_qry_dtl")
	for rows.Next() {
		err := rows.StructScan(&d)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", d)
	}
}
