package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	workers := 3
	numConveyor(workers, arr)
}

func numConveyor(workers int, arr []int) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			writeWorker(ch1, ch2)
		}()
	}

	go func() {
		defer wg.Done()
		readWorker(ch2)
	}()

	go func() {
		for _, item := range arr {
			ch1 <- item
		}
		close(ch1)
	}()

	wg.Wait()
	close(ch2)
}

func writeWorker(ch1, ch2 chan int) {
	for num := range ch1 {
		ch2 <- num * num
	}
}

func readWorker(ch2 chan int) {
	for num := range ch2 {
		fmt.Println(num)
	}
}
