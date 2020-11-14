package tasks

import (
	"database/sql"
	"math"
	"reflect"
	"sync"
)

type Job struct {
	id      int
	connect *sql.DB
	data    interface{}
}

type Result struct {
	job    Job
	status bool
}

var jobs = make(chan Job, 100)
var results = make(chan Result, 100)

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, writeCH(job.connect, job.data)}
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

func allocate(connect *sql.DB, data interface{}, page int) {
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Slice {
		total := v.Len()

		if total <= page {
			job := Job{0, connect, v.Slice(0, total).Interface()}
			jobs <- job
		} else {
			i := 0
			for ; i < total/page; i++ {
				job := Job{i, connect, v.Slice(i*page, (i+1)*page).Interface()}
				jobs <- job
			}

			k := float64(total) / float64(page)
			noOfJobs := int(math.Ceil(k))
			if noOfJobs > total/page {
				job := Job{i, connect, v.Slice((i * page), total).Interface()}
				jobs <- job
			}
		}

		close(jobs)
	}
}

func finish(done chan bool) {
	for _ = range results {
	}
	done <- true
}

func Start(connect *sql.DB, data interface{}) {
	go allocate(connect, data, 2000)
	done := make(chan bool)
	go finish(done)
	noOfWorkers := 100
	createWorkerPool(noOfWorkers)
	<-done
}
