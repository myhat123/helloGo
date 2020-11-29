package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mailru/go-clickhouse"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"

	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
)

var PGURL = "postgresql://jxyz:1234@localhost/jr"

// var CHURL = "tcp://127.0.0.1:9000?username=hzg&password=1234&database=finance"
var CHURL = "http://hzg:1234@localhost:8123/finance"

func GetPG() *pgx.Conn {
	config, err := pgx.ParseConfig(PGURL)

	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	conn.ConnInfo().RegisterDataType(pgtype.DataType{
		Value: &shopspring.Numeric{},
		Name:  "numeric",
		OID:   pgtype.NumericOID,
	})

	return conn
}

func GetCH() *sql.DB {
	connect, _ := sql.Open("clickhouse", CHURL)

	return connect
}
