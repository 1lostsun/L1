package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	incMutex := IncMutex{count: 0, mx: &sync.Mutex{}}
	incMutex.incrementByMutex(10, 34500178)

	incAtomic := IncAtomic{count: atomic.Int64{}}
	incAtomic.incrementByAtomic(10, 34500178)

	fmt.Println(incMutex.count)
	fmt.Println(incAtomic.count.Load())
}

// IncMutex 1 вариант через Mutex
type IncMutex struct {
	count int64
	mx    *sync.Mutex
}

func (inc *IncMutex) incrementByMutex(workers, target int) {
	tasks := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for range tasks {
				inc.incrementWorker()
			}
		}()
	}

	for i := 0; i < target; i++ {
		tasks <- struct{}{}
	}
	close(tasks)
	wg.Wait()
}

func (inc *IncMutex) incrementWorker() {
	inc.mx.Lock()
	defer inc.mx.Unlock()
	inc.count++
}

// IncAtomic 2 вариант через Atomic int
type IncAtomic struct {
	count atomic.Int64
}

func (inc *IncAtomic) incrementByAtomic(workers, target int) {
	tasks := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for range tasks {
				inc.count.Add(1)
			}
		}()
	}

	for i := 0; i < target; i++ {
		tasks <- struct{}{}
	}
	close(tasks)
	wg.Wait()
}
