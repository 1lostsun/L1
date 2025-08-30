package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	sleepWithIterations(2 * time.Second)
	sleepWithChannel(2 * time.Second)
	fmt.Println(time.Now())
}

// 1 вариант через канал
func sleepWithChannel(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

// 2 вариант через ожидание в цикле
func sleepWithIterations(duration time.Duration) {
	end := time.Now().Add(duration)
	for time.Now().Before(end) {
	}
}
