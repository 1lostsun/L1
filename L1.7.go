package main

import (
	"fmt"
	"sync"
)

func main() {
	concurrentMapWritingByMutex(5, 10000)
	concurrentMapWritingBySyncMap(5, 10000)
}

// 1 вариант через mutex
func concurrentMapWritingByMutex(workers, maxRange int) {
	var mx sync.RWMutex
	mp := make(map[string]int)
	wg := sync.WaitGroup{}

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			mapWorker(mp, &mx, i, maxRange)
		}()
	}

	wg.Wait()
	fmt.Println(mp)
}

func mapWorker(mp map[string]int, mx *sync.RWMutex, workerID, maxRange int) {
	for i := 0; i < maxRange; i++ {
		mx.Lock()
		fmt.Println("worker: ", workerID, " is writing")
		mp["mutex"] += 1
		mx.Unlock()
	}
}

// 2 вариант через sync.Map
func concurrentMapWritingBySyncMap(workers, maxRange int) {
	var mp sync.Map
	mp.Store("mutex", 0)
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < maxRange; j++ {
				for {
					count, _ := mp.Load("mutex")
					if mp.CompareAndSwap("mutex", count, count.(int)+1) {
						break
					}
				}
			}
		}()
	}

	wg.Wait()
	fmt.Println(mp.Load("mutex"))
}
