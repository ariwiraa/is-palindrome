package palindrome

import (
	"strings"
)

func IsPalindrome(value string) bool {
	value = strings.ToLower(value)
	lastIndex := len(value) - 1
	for i := 0; i < lastIndex/2 && i < (lastIndex-i); i++ {
		if value[i] != value[lastIndex-i] {
			return false
		}
	}
	return true
}
