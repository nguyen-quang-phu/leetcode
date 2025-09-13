package leetcode

// @leet start
const VOWELS = "AEIOUaeiou"
const MASK= 0b11111
const VOWELMASK =
	(1 << ('A' & MASK)) |
	(1 << ('E' & MASK)) |
	(1 << ('I' & MASK)) |
	(1 << ('O' & MASK)) |
	(1 << ('U' & MASK))

func isVowel(b byte) bool {
	return (VOWELMASK>>(b&MASK))&1 > 0
}

func sortVowels(s string) string {
	freq := ['z' + 1]int{}

	for _, c := range s {
		freq[c]++
	}

	vowelIdx := 0
	t := []byte(s)

	for i := range s {
		if !isVowel(s[i]) {
			continue
		}
		for freq[VOWELS[vowelIdx]] == 0 {
			vowelIdx++
		}
		t[i] = VOWELS[vowelIdx]
		freq[VOWELS[vowelIdx]]--
	}

	return string(t)
}

// @leet end

// Keynold
