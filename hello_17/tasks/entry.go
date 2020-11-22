package tasks

import (
	"github.com/jmoiron/sqlx"

	"hello_17/common"
)

//写入数据至postgres
func writeCH(connect *sqlx.DB, data interface{}) {

	if qrydtl, ok := data.([]*common.DBrchQryDtl); ok {
		insAction(connect, qrydtl)
	}

}
