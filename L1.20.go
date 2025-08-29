package main

import "fmt"

func main() {
	fmt.Println(wordsReverser([]rune("snow dog sun")))
}

func runeReverser(runes []rune, left int, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

func wordsReverser(runes []rune) string {
	runeReverser(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			runeReverser(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}
