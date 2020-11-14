package tasks

import (
	"database/sql"

	"hello_15/common"
)

//写入数据至clickhouse
func writeCH(connect *sql.DB, data interface{}) bool {

	if qrydtl, ok := data.([]*common.DBrchQryDtl); ok {
		insAction(connect, qrydtl)
	}

	return true
}
