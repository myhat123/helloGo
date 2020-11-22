package db

import (
	// "database/sql"
	"log"

	"hello_17/common"
	"hello_17/tasks"
)

func GetBrchQryDtl() {
	db := GetPG()

	rows, err := db.Queryx("SELECT * FROM brch_qry_dtl")

	if err != nil {
		log.Fatalln(err)
	}

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
