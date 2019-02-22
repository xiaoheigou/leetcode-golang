package main

import "fmt"

func main() {
	s := longestPalindrome("aa")
	fmt.Println(s)
}
func longestPalindrome(s string) string {

	if len(s) == 0 {
		return s
	}

	slen := len(s)
	i := 0
	j := 0
	subStr := ""

	for i = slen; i >= 1; i-- {
		highIdx := i - 1
		for j = 0; highIdx < slen; {
			subStr = subString(s, j, highIdx+1)
			if isPalindrome(subStr) == true {
				return subStr
			}
			j++
			highIdx = j + i - 1
		}
	}

	return subStr
}

func subString(s string, low int, high int) string {
	//strbytes := string([]rune(s)[low:high])
	str := s[low:high]
	return str
}

func isPalindrome(s string) bool {
	flag := true
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			flag = false
			return flag
		}
	}

	return flag
}
