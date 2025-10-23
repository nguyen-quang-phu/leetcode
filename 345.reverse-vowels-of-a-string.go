package leetcode

import "fmt"

// @leet start
func isVowel(b byte) bool {
	if b < 'A' || (b > 'Z' && b < 'a') || b > 'z' {
		return false
	}

	const VOWELS = "AEIOUaeiou"
	const MASK= 0b11111
	const VOWELMASK =
		(1 << ('A' & MASK)) |
		(1 << ('E' & MASK)) |
		(1 << ('I' & MASK)) |
		(1 << ('O' & MASK)) |
		(1 << ('U' & MASK))

	return (VOWELMASK>>(b&MASK))&1 > 0
}

func reverseVowels(s string) string {
	left:= 0
	right:= len(s)-1
	leftReady:= false
	rightReady:= false
	result:= []byte(s)

	for left <= right {
		if !leftReady {
			if isVowel(s[left]) {
				leftReady= true
			}
			left++
		}
		if !rightReady {
			if isVowel(s[right]) {
				rightReady= true
			}
			right--
		}

		if leftReady && rightReady {
			result[left-1], result[right+1]= result[right+1], result[left-1]
			leftReady, rightReady= false, false
		}
	}

	return string(result)
}
// @leet end

// Keynold
