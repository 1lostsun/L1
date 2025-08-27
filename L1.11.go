package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 3, 4}
	arraysIntersection(arr1, arr2)
}

func arraysIntersection(arr1, arr2 []int) {
	secondArrToMap := make(map[int]bool)
	var intersection []int

	for _, v := range arr2 {
		secondArrToMap[v] = true
	}

	for _, val := range arr1 {
		if _, ok := secondArrToMap[val]; ok {
			intersection = append(intersection, val)
		}
	}
	fmt.Println(intersection)
}
