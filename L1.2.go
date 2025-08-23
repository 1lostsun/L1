package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	arr := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	const workers = 2

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			Worker(ch)
		}()
	}

	for _, value := range arr {
		ch <- value
	}

	close(ch)
	wg.Wait()
}

func Worker(ch chan int) {
	for data := range ch {
		fmt.Println("some data: ", data)
	}
}
