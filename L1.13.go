package main

import "fmt"

func main() {
	arithmeticSwap(4, 5)
	xorSwap(3, 5)
}

// 1 вариант через сложение/вычитание
func arithmeticSwap(num1, num2 int) {
	num1 -= num2
	num2 += num1
	num1 = num2 - num1

	fmt.Println(num1, num2)
}

// 2 вариант через XOR-обмен
func xorSwap(num1, num2 int) {
	num1 ^= num2
	num2 ^= num1
	num1 ^= num2

	fmt.Println(num1, num2)
}
