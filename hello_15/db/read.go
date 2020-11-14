package db

import (
	"context"
	// "fmt"
	"log"

	"hello_15/common"
	"hello_15/tasks"
)

func GetBrchQryDtl() {

	dbpool := GetPG()
	connect := GetCH()

	defer dbpool.Close()
	defer connect.Close()

	rows, _ := dbpool.Query(context.Background(), `
		select acc, tran_date, amt, dr_cr_flag, rpt_sum, timestamp1 from brch_qry_dtl
	`)

	records := make([]*common.DBrchQryDtl, 0)
	i := 0
	for rows.Next() {
		d := new(common.DBrchQryDtl)

		err := rows.Scan(&d.Acc, &d.TranDate, &d.Amt, &d.DrCrFlag, &d.RptSum, &d.Timestamp1)
		if err != nil {
			log.Fatal("Scan failed", err)
		}

		if i > 0 && i%(100*10000) == 0 {
			tasks.Start(connect, records)
			records = make([]*common.DBrchQryDtl, 0)
		}

		records = append(records, d)

		i = i + 1
	}

	// for _, t := range records {
	// 	fmt.Println(*t)
	// }

	tasks.Start(connect, records)
}
