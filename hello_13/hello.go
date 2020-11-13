package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"

	"sync"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"

	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
)

type DBrchQryDtl struct {
	Acc        string
	TranDate   time.Time
	Amt        decimal.Decimal
	DrCrFlag   int
	RptSum     string
	Timestamp1 string
}

type Job struct {
	id      int
	connect *sql.DB
	data    []*DBrchQryDtl
}

type Result struct {
	job    Job
	status bool
}

var jobs = make(chan Job, 100)
var results = make(chan Result, 100)

func writeCH(job Job) bool {
	// for _, t := range records {
	// 	fmt.Println(*t)
	// }
	// fmt.Println(job.id, len(job.data))

	tx, _ := job.connect.Begin()
	stmt, _ := tx.Prepare(`
		INSERT INTO brch_qry_dtl (tran_date, timestamp1, acc, amt, dr_cr_flag, rpt_sum) 
		VALUES (?, ?, ?, ?, ?, ?)
	`)

	for _, t := range job.data {
		fmt.Println(t.TranDate, t.Timestamp1, t.Acc, t.Amt, t.DrCrFlag, t.RptSum)
		x, _ := t.Amt.Float64()

		if _, err := stmt.Exec(t.TranDate, t.Timestamp1, t.Acc, x, t.DrCrFlag, t.RptSum); err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()
	return true
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, writeCH(job)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(connect *sql.DB, data []*DBrchQryDtl, page int) {
	total := len(data)
	if total <= page {
		job := Job{0, connect, data[:total]}
		jobs <- job
	} else {
		i := 0
		for ; i < total/page; i++ {
			job := Job{i, connect, data[i*page : (i+1)*page]}
			jobs <- job
		}

		k := float64(total) / float64(page)
		noOfJobs := int(math.Ceil(k))
		if noOfJobs > total/page {
			job := Job{i, connect, data[i*page : total]}
			jobs <- job
		}
	}
	close(jobs)
}

func finish(done chan bool) {
	for _ = range results {
	}
	done <- true
}

func Start(connect *sql.DB, data []*DBrchQryDtl) {
	go allocate(connect, data, 2000)
	done := make(chan bool)
	go finish(done)
	noOfWorkers := 100
	createWorkerPool(noOfWorkers)
	<-done
}

func main() {

	config, err := pgxpool.ParseConfig("postgresql://jxyz:1234@localhost/jr")
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
	// defer conn.Close(context.Background())

	defer dbpool.Close()

	connect, _ := sql.Open("clickhouse", "tcp://127.0.0.1:9000?username=hzg&password=1234&database=finance&debug=true")

	// var acc string
	// var amt decimal.Decimal

	// err = conn.QueryRow(context.Background(), `
	// 	select acc, amt from brch_qry_dtl
	// 	where tran_date=$1
	// `, "2019-11-27").Scan(&acc, &amt)

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(acc, amt)

	rows, _ := dbpool.Query(context.Background(), `
		select acc, tran_date, amt, dr_cr_flag, rpt_sum, timestamp1 from brch_qry_dtl
	`)

	records := make([]*DBrchQryDtl, 0)
	i := 0
	for rows.Next() {
		d := new(DBrchQryDtl)

		err := rows.Scan(&d.Acc, &d.TranDate, &d.Amt, &d.DrCrFlag, &d.RptSum, &d.Timestamp1)
		if err != nil {
			log.Fatal("Scan failed", err)
		}

		if i > 0 && i%(100*10000) == 0 {
			Start(connect, records)
			records = make([]*DBrchQryDtl, 0)
		}

		records = append(records, d)

		i = i + 1
	}

	// for _, t := range records {
	// 	fmt.Println(*t)
	// }

	Start(connect, records)
}
