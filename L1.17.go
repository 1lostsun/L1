package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 3
	indexByRecursion := binarySearchByRecursion(nums, target, 0, len(nums))
	indexByIterations := binarySearchByIterations(nums, target)
	fmt.Println(indexByRecursion, indexByIterations)
}

// 1 вариант через рекурсию
func binarySearchByRecursion(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binarySearchByRecursion(nums, target, left, mid-1)
	} else {
		return binarySearchByRecursion(nums, target, mid+1, right)
	}
}

// 2 вариант через итерации
func binarySearchByIterations(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
