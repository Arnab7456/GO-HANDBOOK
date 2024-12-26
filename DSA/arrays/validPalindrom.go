package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		leftChar := s[left]
		rightChar := s[right]

		if !unicode.IsLetter(rune(leftChar)) && !unicode.IsDigit(rune(leftChar)){
		 left++
		}else if !unicode.IsLetter(rune(rightChar)) && !unicode.IsDigit(rune(rightChar)){
			right--
		} else {
			if strings.ToLower(string(leftChar)) != strings.ToLower(string(rightChar)){
				return false
			}
			left++
		right--
		}
		
	}
	return true
}	