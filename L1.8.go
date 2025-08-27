package main

import "fmt"

func main() {
	fmt.Println(bitSwap(5, 0, 0))  // 0101 -> 0100
	fmt.Println(bitSwap(8, 0, 1))  // 01000 -> 01001
	fmt.Println(bitSwap(64, 7, 1)) // 01000000 -> 11000000
	fmt.Println(bitSwap(96, 5, 0)) // 1100000 -> 1000000
}

func bitSwap(example, x, binary int64) int64 {
	var mask int64 = 1 << x
	if binary == 0 {
		return example ^ mask
	}
	return example | mask
}
