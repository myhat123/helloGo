package tasks

import (
	"database/sql"
	"log"

	"hello_18/common"
)

func insAction(connect *sql.DB, records []*common.DBrchQryDtl) {
	tx, _ := connect.Begin()
	stmt, _ := tx.Prepare(`
		INSERT INTO brch_qry_dtl (tran_date, timestamp1, acc, amt, dr_cr_flag, rpt_sum) 
		VALUES (?, ?, ?, ?, ?, ?)
	`)

	for _, t := range records {
		x := t.Amt.String()

		if _, err := stmt.Exec(t.TranDate, t.Timestamp1, t.Acc, x, t.DrCrFlag, t.RptSum); err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()
}
