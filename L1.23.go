package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(deleteElemByIndex(arr, 4))
}

func deleteElemByIndex(arr []int, index int) []int {
	_ = copy(arr[index:], arr[index+1:])
	arr = arr[:len(arr)-1]

	return arr
}
