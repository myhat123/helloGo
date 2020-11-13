package tasks

import (
	"log"
)

func writeCH(job Job) bool {
	log.Printf("job.id: %d, data length: %d", job.id, len(job.data))

	tx, _ := job.connect.Begin()
	stmt, _ := tx.Prepare(`
		INSERT INTO brch_qry_dtl (tran_date, timestamp1, acc, amt, dr_cr_flag, rpt_sum) 
		VALUES (?, ?, ?, ?, ?, ?)
	`)

	for _, t := range job.data {
		x, _ := t.Amt.Float64()

		if _, err := stmt.Exec(t.TranDate, t.Timestamp1, t.Acc, x, t.DrCrFlag, t.RptSum); err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()
	return true
}
