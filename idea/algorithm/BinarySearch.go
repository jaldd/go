package main

import "fmt"

func BinarySearch(array []int, target, l, r int) int {

	if l > r {
		return -1
	}
	mid := (l + r + 1) / 2
	midVal := array[mid]

	if midVal == target {
		return mid
	} else if midVal > target {
		return BinarySearch(array, target, 0, mid-1)
	} else {
		return BinarySearch(array, target, mid+1, r)
	}
}

func main() {

	array := []int{1, 5, 9, 15, 81, 89123, 189, 333}
	target := 500
	result := BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)

	target = 189
	result = BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)
}
