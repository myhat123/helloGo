package tasks

import (
	"database/sql"

	_ "github.com/mailru/go-clickhouse"

	"hello_18/common"
)

func GetCH() *sql.DB {
	connect, _ := sql.Open("clickhouse", common.CHURL)

	return connect
}
