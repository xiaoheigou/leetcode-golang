package main

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func TestTwoSum(t *testing.T) {
	arr := []int{2, 3, 4, 5, 6, 7, 8}
	result := twoSum(arr, 15)
	if !Equal(result, []int{5, 6}) {
		t.Errorf("wrong suppose to be 5,6,accully is %v", result)
	}
}
