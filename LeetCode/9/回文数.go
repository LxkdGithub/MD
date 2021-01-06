package main

import "strconv"

func main() {
	println(isPalindrome(-23))
	println(isPalindrome(5))
	println(isPalindrome(45654))
	println(isPalindrome(34465))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	for i:=0;i<=len(s)/2;i++ {
		if s[i] != s[len(s) - i-1] {
			return false
		}
	}
	return true
}
