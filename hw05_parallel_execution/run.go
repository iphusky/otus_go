package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n int, m int) error {
	if n <= 0 || len(tasks) == 0 {
		return nil
	}

	if m < 0 {
		m = 0
	}

	errorCount := int32(m)

	tasksQueue := make(chan Task, len(tasks))
	runProducer(tasksQueue, tasks)
	runConsumer(tasksQueue, n, &errorCount)

	if atomic.LoadInt32(&errorCount) >= 0 {
		return nil
	}

	return ErrErrorsLimitExceeded
}

func runProducer(tasksQueue chan<- Task, tasks []Task) {
	defer close(tasksQueue)

	for _, task := range tasks {
		tasksQueue <- task
	}
}

func runConsumer(tasksQueue <-chan Task, processCount int, errorCount *int32) {
	var wg sync.WaitGroup
	wg.Add(processCount)

	for i := 0; i < processCount; i++ {
		go runProcess(tasksQueue, errorCount, &wg)
	}

	wg.Wait()
}

func runProcess(tasksQueue <-chan Task, errorCount *int32, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if task, ok := <-tasksQueue; ok && atomic.LoadInt32(errorCount) >= 0 {
			execute(task, errorCount)
		} else {
			return
		}
	}
}

func execute(task Task, errorCount *int32) {
	if err := task(); err != nil {
		atomic.AddInt32(errorCount, -1)
	}
}
