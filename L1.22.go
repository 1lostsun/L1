package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(100), big.NewInt(200)
	fmt.Println(bigMultiplication(a, b))
}

func bigMultiplication(a, b *big.Int) *big.Int {
	var mul big.Int
	mul.Mul(a, b)
	return &mul
}
