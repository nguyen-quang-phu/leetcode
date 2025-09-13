package leetcode

// @leet start

func isVowel(b byte) bool {
	const VOWELS = "AEIOUaeiou"
	const MASK = 0b11111
	const VOWELMASK = (1 << ('A' & MASK)) |
		(1 << ('E' & MASK)) |
		(1 << ('I' & MASK)) |
		(1 << ('O' & MASK)) |
		(1 << ('U' & MASK))

	return (VOWELMASK>>(b&MASK))&1 > 0
}

func maxFreqSum(s string) int {
	freq := ['z' + 1]int{}
	for i := range s {
		freq[s[i]]++
	}
	maxVowelFreq := 0
	maxConsonantFreq := 0

	for i := range freq {
		if isVowel(byte(i)) {
			if freq[i] > maxVowelFreq {
				maxVowelFreq = freq[i]
			}
		} else {
			if freq[i] > maxConsonantFreq {
				maxConsonantFreq = freq[i]
			}
		}
	}
	return maxVowelFreq + maxConsonantFreq
}

// @leet end

// Keynold
