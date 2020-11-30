package tasks

import (
	"hello_18/common"
)

//写入数据至clickhouse
func writeCH(data interface{}) {

	if qrydtl, ok := data.([]*common.DBrchQryDtl); ok {
		insAction(qrydtl)
	}

}
