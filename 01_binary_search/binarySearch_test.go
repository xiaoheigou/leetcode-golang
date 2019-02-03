package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testArray := []int{1, 2, 3, 4, 5, 6, 7, 77, 665}
	num := BinarySearch(testArray, 1)

	if num == -1 {
		t.Log("not found")
	} else {
		t.Log(num)
	}
	testArray2 := []int{21, 22, 33, 54, 55, 66, 87, 747, 4665}
	num2 := BinarySearch(testArray2, 55)

	if num2 == -1 {
		t.Log("not found")
	} else {
		t.Log(num2)
	}

}
