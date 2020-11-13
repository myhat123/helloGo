package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
)

var PGURL = "postgresql://jxyz:1234@localhost/jr"
var CHURL = "tcp://127.0.0.1:9000?username=hzg&password=1234&database=finance"

func GetPG() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(PGURL)
	if err != nil {
		log.Fatal(err)
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &shopspring.Numeric{},
			Name:  "numeric",
			OID:   pgtype.NumericOID,
		})

		return nil
	}

	dbpool, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	// defer dbpool.Close()

	return dbpool
}

func GetCH() *sql.DB {
	connect, _ := sql.Open("clickhouse", CHURL)

	return connect
}
