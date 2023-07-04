package worker_pool

import (
	"fmt"
	"log"
	"sync"
)

type Pool interface {
	// Start gets worker-pool ready to process jobs and should only be called once
	Start()
	// Stop stops worker-pool and should only be called once
	Stop()
	// AddWork adds task for worker pool to process. Only valid after Start and before Stop
	AddWork(Task)
}

type Task interface {
	// Execute Process the task
	Execute() error
	// OnFailure Handle any error returned by Execute
	OnFailure(error)
}

type SimplePool struct {
	numOfWorkers int
	tasks        chan Task
	start        sync.Once     // make sure pool start once
	stop         sync.Once     // make sure pool stop once
	quit         chan struct{} // signal to stop working
}

var ErrNoWorkers = fmt.Errorf("attempting to create worker pool with less than 1 worker")
var ErrNegativeChannelSize = fmt.Errorf("attempting to create worker pool with a negative channel size")

func NewSimplePool(numWorkers int, channelSize int) (Pool, error) {
	if numWorkers <= 0 {
		return nil, ErrNoWorkers
	}
	if channelSize < 0 {
		return nil, ErrNegativeChannelSize
	}
	tasks := make(chan Task, channelSize)
	return &SimplePool{
		numOfWorkers: numWorkers,
		tasks:        tasks,
		start:        sync.Once{},
		stop:         sync.Once{},
		quit:         make(chan struct{}),
	}, nil
}

func (pool *SimplePool) Start() {
	pool.start.Do(func() {
		log.Print("[starting] worker pool")
		pool.startWorkers()
	})
}

func (pool *SimplePool) Stop() {
	pool.stop.Do(func() {
		log.Print("[stopping] worker pool")
		close(pool.quit)
	})
}

// AddWork - If channel buffer is empty or full or all workers busy, process will hang until work consumed/stop is called
func (pool *SimplePool) AddWork(t Task) {
	select {
	case pool.tasks <- t:
	case <-pool.quit:
	}
}

func (pool *SimplePool) startWorkers() {
	for i := 0; i < pool.numOfWorkers; i++ {
		go func(workerNum int) {
			log.Printf("[starting] worker : %d", workerNum)
			for {
				select {
				case <-pool.quit:
					log.Printf("[stopping] worker : %d with quit signal\n", workerNum)
					return
				case task, ok := <-pool.tasks:
					if !ok {
						log.Printf("[stopping] worker : %d with closed task channel\n", workerNum)
						return
					}
					if err := task.Execute(); err != nil {
						task.OnFailure(err)
					}
				}
			}
		}(i)
	}
}
