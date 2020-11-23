package tasks

import (
	"github.com/jmoiron/sqlx"

	"hello_17/common"
)

//写入数据至postgres
func writePG(db *sqlx.DB, data interface{}) {

	switch v := data.(type) {
	case []*common.DBrchQryDtl:
		insAction(db, v)
	}
}
