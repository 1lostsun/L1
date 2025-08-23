package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	workerStart(5)
}

func worker(ch chan int) {
	for data := range ch {
		fmt.Println("some data: ", data)
	}
}

func workerStart(workers int) {
	ticker := time.NewTicker(500 * time.Millisecond)
	ch := make(chan int, 5)
	wg := sync.WaitGroup{}

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			worker(ch)
		}()
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-stop:
				close(ch)
			case <-ticker.C:
				ch <- rand.Intn(1000) * rand.Intn(100)
			}
		}
	}()

	wg.Wait()
}
