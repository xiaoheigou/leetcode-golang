package main

func BinarySearch(array []int, target int) int {
	return BinarySearchMux(array, 0, len(array)-1, target)
}
func BinarySearchMux(array []int, min int, max int, target int) int {
	if min > max {
		return -1
	} else {
		mid := (min + max) / 2
		if array[mid] > target {
			return BinarySearchMux(array, min, mid-1, target)
		} else if array[mid] < target {
			return BinarySearchMux(array, mid+1, max, target)
		} else {
			return mid
		}
	}
}
