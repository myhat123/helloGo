package tasks

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"math"
	"reflect"
	"sync"

	"github.com/panjf2000/ants/v2"
)

type Job struct {
	id   int
	db   *sqlx.DB
	data interface{}
}

var jobs = make(chan Job, 100)

func allocate(db *sqlx.DB, data interface{}, page int) {
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Slice {
		total := v.Len()

		if total <= page {
			job := Job{0, db, v.Slice(0, total).Interface()}
			jobs <- job
		} else {
			i := 0
			for ; i < total/page; i++ {
				job := Job{i, db, v.Slice(i*page, (i+1)*page).Interface()}
				jobs <- job
			}

			k := float64(total) / float64(page)
			noOfJobs := int(math.Ceil(k))
			if noOfJobs > total/page {
				job := Job{i, db, v.Slice((i * page), total).Interface()}
				jobs <- job
			}
		}

		close(jobs)
	}
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup

	p, _ := ants.NewPoolWithFunc(noOfWorkers, func(job interface{}) {
		if j, ok := job.(Job); ok {
			writeCH(j.db, j.data)
		}
		wg.Done()
	})

	for job := range jobs {
		wg.Add(1)
		_ = p.Invoke(job)
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
}

func Start(db *sqlx.DB, data interface{}) {
	go allocate(db, data, 20)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
}
