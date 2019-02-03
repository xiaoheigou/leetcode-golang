package main

func lengthOfLongestSubstring(s string) int {
	lastOccuerd := make(map[rune]int)

	start, length := 0, 0
	for i, v := range s {
		if last, ok := lastOccuerd[v]; ok && last >= start {
			start = last + 1
		}
		if i-start+1 > length {
			length = i - start + 1
		}
		lastOccuerd[v] = i
	}
	return length
}
