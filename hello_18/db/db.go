package db

import (
	"context"
	"log"

	_ "github.com/mailru/go-clickhouse"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"

	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"

	"hello_18/common"
)

func GetPG() *pgx.Conn {
	config, err := pgx.ParseConfig(common.PGURL)

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
