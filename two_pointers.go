package twopointers

import (
	"regexp"
	"strings"
)

// IsPalindrome returns true if a given string is a palindrome or false if it's not
func IsPalindrome(str string) bool {
	charArr := transformString(str)
	p1 := 0
	p2 := len(charArr) - 1
	for p1 < p2 {
		if charArr[p1] != charArr[p2] {
			return false
		}
		p1++
		p2--
	}
	return true
}

func transformString(str string) []rune {
	resp := strings.ToLower(str)
	resp = clearNonAlphanumeric(resp)
	return []rune(resp)
}

func clearNonAlphanumeric(str string) string {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-z0-9]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
