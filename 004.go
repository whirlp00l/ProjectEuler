package main

import "fmt"
import "strconv"

func isPalindrome(num int) (res bool) {
	s := strconv.Itoa(num)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	res, n := 0, 0
	for a := 999; a > 800; a-- {
		for b := 999; b > 800; b-- {
			n = a * b
			if isPalindrome(n) {
				if res < n {
					res = n
				}
			}
		}
	}
	fmt.Println(res)
}
