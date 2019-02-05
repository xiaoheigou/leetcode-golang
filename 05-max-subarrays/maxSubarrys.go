package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{
		3, -2, 34, 3, -3, -23, -54, 23, -2, 23, 2, -2, -2, 3,
	}
	l, h, s := findMaxSubarray(a, 0, len(a)-1)
	fmt.Println(l, h, s)

}
func findMaxSubarray(a []int, low, high int) (arrLow, arrHigh, arrSum int) {
	if low == high {
		return low, low, a[low]
	}
	mid := (low + high) / 2
	//	fmt.Println(low, mid, high)
	leftLow, leftHigh, leftSum := findMaxSubarray(a, low, mid)

	rightLow, rightHigh, rightSum := findMaxSubarray(a, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossSubarray(a, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}
	return crossLow, crossHigh, crossSum
}
func findMaxCrossSubarray(a []int, low, mid, high int) (arrLow, arrHigh, arrSum int) {
	leftSum := math.MinInt64
	sum := 0
	for i := mid; i >= low; i-- {
		sum = sum + a[i]
		if sum > leftSum {
			leftSum = sum
			arrLow = i
		}
	}
	rightSum := math.MinInt64
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum = sum + a[i]
		if sum > rightSum {
			rightSum = sum
			arrHigh = i
		}
	}
	return arrLow, arrHigh, leftSum + rightSum

}
