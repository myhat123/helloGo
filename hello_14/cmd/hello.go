package main

import (
	"hello_14/db"
)

func main() {
	dbpool := db.GetPG()
	connect := db.GetCH()

	db.GetBrchQryDtl(dbpool, connect)
}
