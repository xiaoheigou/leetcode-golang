package main

import "strings"

func main() {

}
func romanToInt(s string) int {
	summe := 0
	i := 0
	s = strings.ToLower(s)
	dickey := map[byte]int{'i': 1, 'v': 5, 'x': 10, 'l': 50, 'c': 100, 'd': 500, 'm': 1000}
	for i < len(s) {
		if i+1 < len(s) {
			if dickey[s[i]] < dickey[s[i+1]] {
				summe += dickey[s[i+1]] - dickey[s[i]]
				i += 2
				continue
			} else {
				summe += dickey[s[i]]
			}
		} else {
			summe += dickey[s[i]]

		}
		i += 1
	}
	return summe

}
