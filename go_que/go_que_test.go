package go_que

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/tnclong/go-que"
	"github.com/tnclong/go-que/pg"
	"github.com/tnclong/go-que/scheduler"
)

func TestA(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://myuser:mypassword@127.0.0.1:5432/mydb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	q, err := pg.New(db)
	if err != nil {
		panic(err)
	}

	// consumer
	qsConsumer := "que.consumer"
	{
		worker, err := que.NewWorker(que.WorkerOptions{
			Queue:              qsConsumer,
			Mutex:              q.Mutex(),
			MaxLockPerSecond:   10,
			MaxBufferJobsCount: 0,

			Perform: func(ctx context.Context, job que.Job) error {
				log.Printf("consume job msg at %v", time.Now())
				return job.Done(ctx)
			},
			MaxPerformPerSecond:       2,
			MaxConcurrentPerformCount: 1,
		})
		if err != nil {
			panic(err)
		}
		defer worker.Stop(context.Background())
		go func() {
			err := worker.Run()
			fmt.Println("Run():", err.Error())
		}()
	}

	// scheduler logic
	{
		qsScheduler := "que.cronjob"
		sc := &scheduler.Scheduler{
			DB:      db,
			Queue:   qsScheduler,
			Enqueue: q.Enqueue,
			Provider: &scheduler.MemProvider{
				Schedule: scheduler.Schedule{
					"cronjob.for." + qsConsumer: scheduler.Item{
						Queue:          qsConsumer,
						Args:           ``,
						Cron:           "* * * * *",
						RecoveryPolicy: scheduler.Ignore,
					},
				},
			},
			Derivations: nil,
		}
		scWorker, err := que.NewWorker(que.WorkerOptions{
			Queue:                     qsScheduler,
			Mutex:                     q.Mutex(),
			MaxLockPerSecond:          10,
			MaxBufferJobsCount:        0,
			Perform:                   sc.Perform,
			MaxPerformPerSecond:       2,
			MaxConcurrentPerformCount: 1,
		})
		if err != nil {
			panic(err)
		}
		defer scWorker.Stop(context.Background())
		go func() {
			err := scWorker.Run()
			fmt.Println("Run():", err.Error())
		}()
		if err = sc.Prepare(context.Background()); err != nil {
			panic(err)
		}
	}

	select {}
}
