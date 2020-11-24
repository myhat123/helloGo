package tasks

import (
	"database/sql"

	"hello_18/common"
)

//写入数据至clickhouse
func writeCH(connect *sql.DB, data interface{}) {

	if qrydtl, ok := data.([]*common.DBrchQryDtl); ok {
		insAction(connect, qrydtl)
	}

}
