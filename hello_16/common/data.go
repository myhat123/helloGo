package common

import (
	"time"

	"github.com/shopspring/decimal"
)

type DBrchQryDtl struct {
	Acc        string
	TranDate   time.Time
	Amt        decimal.Decimal
	DrCrFlag   int
	RptSum     string
	Timestamp1 string
}
