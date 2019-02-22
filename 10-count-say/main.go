package main

import "fmt"

func main() {

}

func countAndSay(n int) string {
	s := "1"
	for i := 0; i < n-1; i++ {
		lt, temp, count := s[0], "", 0
		for j := 0; j < len(s); j++ {
			if lt == s[j] {
				count += 1
			} else {
				temp = fmt.Sprintf("%s%d%c", temp, count, lt)
				lt = s[j]
				count = 1
			}
		}
		temp = fmt.Sprintf("%s%d%c", temp, count, lt)
		s = temp
	}
	return s
}
