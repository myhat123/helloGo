package db

import (
	"context"
	"database/sql"
	// "fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"hello_14/comm"
	"hello_14/tasks"
)

func GetBrchQryDtl(dbpool *pgxpool.Pool, connect *sql.DB) {

	rows, _ := dbpool.Query(context.Background(), `
		select acc, tran_date, amt, dr_cr_flag, rpt_sum, timestamp1 from brch_qry_dtl
	`)

	records := make([]*comm.DBrchQryDtl, 0)
	i := 0
	for rows.Next() {
		d := new(comm.DBrchQryDtl)

		err := rows.Scan(&d.Acc, &d.TranDate, &d.Amt, &d.DrCrFlag, &d.RptSum, &d.Timestamp1)
		if err != nil {
			log.Fatal("Scan failed", err)
		}

		if i > 0 && i%(100*10000) == 0 {
			tasks.Start(connect, records)
			records = make([]*comm.DBrchQryDtl, 0)
		}

		records = append(records, d)

		i = i + 1
	}

	// for _, t := range records {
	// 	fmt.Println(*t)
	// }

	tasks.Start(connect, records)
}
