package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	programWithContext()
	programWithOs()
}

// 1 вариант через контекст
func programWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ticker := time.NewTicker(time.Second)

	go func(ctx context.Context) {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				fmt.Println("context program", rand.Intn(1000)*rand.Intn(100))
			}
		}
	}(ctx)

	<-ctx.Done()
}

// 2 вариант через os.SIGINT
func programWithOs() {
	wg := sync.WaitGroup{}
	ticker := time.NewTicker(time.Second)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		for {
			select {
			case <-stop:
				wg.Done()
			case <-ticker.C:
				fmt.Println("os program", rand.Intn(1000)*rand.Intn(100))
			}
		}
	}()

	wg.Wait()
}
