package tasks

import (
	"database/sql"
	"math"
	"sync"

	"hello_14/comm"
)

type Job struct {
	id      int
	connect *sql.DB
	data    []*comm.DBrchQryDtl
}

type Result struct {
	job    Job
	status bool
}

var jobs = make(chan Job, 100)
var results = make(chan Result, 100)

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

func allocate(connect *sql.DB, data []*comm.DBrchQryDtl, page int) {
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

func Start(connect *sql.DB, data []*comm.DBrchQryDtl) {
	go allocate(connect, data, 2000)
	done := make(chan bool)
	go finish(done)
	noOfWorkers := 100
	createWorkerPool(noOfWorkers)
	<-done
}
