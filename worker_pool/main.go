package worker_pool

import (
	"go-snowflake-service/conf"
	"log"
)

var WorkerPool chan func() error

func Start() {
	for _ = range make([]int, conf.ConfigObject.WorkerPoolSize) {
		go func() {
			select {
			case job := <-WorkerPool:
				err := job()
				if err != nil {
					log.Println("[Error] ", err)
				}
			}
		}()
	}
}

func Submit(job func() error) {
	WorkerPool <- job
}
