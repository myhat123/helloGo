package common

import (
	"time"
)

type DBrchQryDtl struct {
	Acc        string
	TranDate   time.Time `db:"tran_date"`
	Amt        string
	DrCrFlag   int    `db:"dr_cr_flag"`
	RptSum     string `db:"rpt_sum"`
	Timestamp1 string
}
