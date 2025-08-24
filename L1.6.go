package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	standartGoroutineCompletion()
	goroutineCompletionByReturnStatement()
	goroutineCompletionByPanic()
	goroutineCompletionByChannelClose()
	goroutineCompletionByContext()
	goroutineCompletionByOS()
	goroutineCompletionByOsExit()
	goroutineCompletionByRuntimeGoExit()

}

func standartGoroutineCompletion() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello World")
	}()

	wg.Wait()
}

func goroutineCompletionByReturnStatement() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if true {
			fmt.Println("return")
			return
		}
		fmt.Println("this part don't go in stdout")
	}()

	wg.Wait()
}

func goroutineCompletionByPanic() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic:", err)
			}
		}()
		panic("boom")
	}()

	wg.Wait()
}

func goroutineCompletionByChannelClose() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(ch <-chan int) {
		defer wg.Done()
		for item := range ch {
			fmt.Println(item)
		}
		fmt.Println("channel closed, goroutine exits")
	}(ch)

loop:
	for {
		dig := rand.Intn(10)
		if dig == 7 {
			close(ch)
			break loop
		}
		ch <- dig
	}

	wg.Wait()
}

func goroutineCompletionByContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(rand.Intn(10))
			}
		}
	}()

	wg.Wait()
}

func goroutineCompletionByOS() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Second)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				return
			case <-ticker.C:
				fmt.Println(rand.Intn(10))
			}
		}
	}()

	wg.Wait()
}

func goroutineCompletionByOsExit() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			dig := rand.Intn(10)
			if dig == 7 {
				os.Exit(1)
			}
			fmt.Println(dig)
		}
	}()

	wg.Wait()
}

func goroutineCompletionByRuntimeGoExit() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			dig := rand.Intn(10)
			if dig == 7 {
				runtime.Goexit()
			}
			fmt.Println(dig)
		}
	}()

	wg.Wait()
}
