package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		//1.judge target - v in map
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
		//2.if dont find ,save v into map m
	}
	return nil
}
