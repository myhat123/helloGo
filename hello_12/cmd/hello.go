package main

import (
	// "fmt"
	"time"
	"gopkg.in/inf.v0"
)

import "hello_12/db"

//d 是购买
var d1 = db.Purchase{"36015828", time.Date(2017, time.October, 20, 0, 0, 0, 0, time.UTC), "132112", "171020710975085783", "710910010742065", "1300099C89", inf.NewDec(5000000, 2), "603302"}
var d2 = db.Purchase{"36015828", time.Date(2018, time.February, 22, 0, 0, 0, 0, time.UTC), "120642", "180222710986153573", "710910010742065", "1300099C89", inf.NewDec(7000000, 2), "603302"}
var d3 = db.Purchase{"36015828", time.Date(2018, time.February, 28, 0, 0, 0, 0, time.UTC), "125759", "180228710987050869", "710910010742065", "1300099C89", inf.NewDec(5000000, 2), "603302"}

func main() {
	var r db.PurchaseRecords

	r.Append(d2)
	r.Append(d1)
	r.Append(d3)

	r.Print()

	r.Sort()
	
	r.Print()
}