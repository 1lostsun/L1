package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	makeSet(arr)
}

func makeSet(arr []string) {
	mp := make(map[string]int)
	var set []string

	for _, v := range arr {
		mp[v]++
	}

	for k, _ := range mp {
		set = append(set, k)
	}

	fmt.Println(set)
}
