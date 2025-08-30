package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isSet("abcd)🥲"))
}

func isSet(st string) bool {
	mp := make(map[string]int)
	for i := range st {
		symbol := strings.ToLower(string(st[i]))
		mp[symbol]++
		if mp[symbol] > 1 {
			return false
		}
	}
	return true
}
