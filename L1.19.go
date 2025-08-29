package main

import "fmt"

func main() {
	fmt.Println(stringReverser("⊃∔±÷∉₶😅😅😅😅🥲{)(}"))
}

func stringReverser(str string) string {
	runeArr := []rune(str)
	var newStr string

	for i := len(runeArr) - 1; i >= 0; i-- {
		newStr += string(runeArr[i])
	}

	return newStr
}
