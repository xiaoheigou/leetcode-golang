package main

import "testing"

func TestLengOfLongestSubstring(t *testing.T) {
	if lengthOfLongestSubstring("abcaabbdgk") != 4 {
		t.Errorf("not true,suppose to be 4 ,accurely is %v", lengthOfLongestSubstring("abcaabbdgsk"))
	}
}
