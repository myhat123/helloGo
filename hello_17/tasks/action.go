package tasks

import (
	"github.com/jmoiron/sqlx"

	"hello_17/common"
)

func insAction(db *sqlx.DB, records []*common.DBrchQryDtl) {
	tx := db.MustBegin()

	for _, t := range records {
		tx.MustExec(`
			INSERT INTO brch_qry_dtl_02 (tran_date, timestamp1, acc, amt, dr_cr_flag, rpt_sum) 
			VALUES ($1, $2, $3, $4, $5, $6)
		`, t.TranDate, t.Timestamp1, t.Acc, t.Amt, t.DrCrFlag, t.RptSum)
	}

	tx.Commit()
}
