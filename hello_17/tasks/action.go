package tasks

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"hello_17/common"
)

func insAction(db *sqlx.DB, records []*common.DBrchQryDtl) {
	tx := db.MustBegin()

	for _, t := range records {

		ds := goqu.Insert("brch_qry_dtl_02").Rows(t)
		insertSQL, _, _ := ds.ToSQL()

		tx.MustExec(insertSQL)
	}

	tx.Commit()
}
