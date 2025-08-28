package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 3, 2, 8, 6, 12, 9}
	sortedByFirst := quickSortFirstElem(arr)
	sortedByMiddle := quickSortMiddleElem(arr)
	fmt.Println(sortedByFirst, sortedByMiddle)
}

// 1 вариант с опорным элементом в начале массива (первый элемент)
func quickSortFirstElem(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left []int
	var right []int

	for _, num := range arr[1:] {
		if num <= pivot {
			left = append(left, num)
			continue
		}
		right = append(right, num)
	}

	return append(append(quickSortFirstElem(left), pivot), quickSortFirstElem(right)...)
}

// 2 вариант с опорным элементом в середине
func quickSortMiddleElem(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left []int
	var right []int

	for _, num := range arr {
		if num == pivot {
			continue
		}
		if num <= pivot {
			left = append(left, num)
			continue
		}
		right = append(right, num)
	}

	return append(append(quickSortMiddleElem(left), pivot), quickSortMiddleElem(right)...)
}
