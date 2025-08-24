package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	timeout := 3 * time.Second
	programWithTimeout(timeout)
}

func programWithTimeout(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	ch := make(chan any)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for item := range ch {
			fmt.Println(item)
		}
	}()

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-ticker.C:
			ch <- rand.Intn(1000)
		}
	}

	close(ch)
	wg.Wait()
}
