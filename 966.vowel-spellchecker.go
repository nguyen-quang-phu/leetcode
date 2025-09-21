package leetcode

import "strings"

// @leet start

func isVowel(b rune) bool {
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

func spellchecker(wordlist []string, queries []string) []string {
	// exact match
	exact := map[string]bool{}
	ci := map[string]string{}
	d := map[string]string{}
	for _, w := range wordlist {
		exact[w] = true
		k := strings.ToLower(w)
		if _, ok := ci[k]; !ok {
			ci[k] = w
		}
		k = devowel(w)
		if _, ok := d[k]; !ok {
			d[k] = w
		}
	}

	r := make([]string, len(queries))
	for i, q := range queries {
		if _, ok := exact[q]; ok {
			r[i] = q
			continue
		}
		if orig, ok := ci[strings.ToLower(q)]; ok {
			r[i] = orig
			continue
		}
		if orig, ok := d[devowel(q)]; ok {
			r[i] = orig
			continue
		}
		r[i] = ""
	}

	return r
}

func devowel(word string) string {
	lower := strings.ToLower(word)
	var sb strings.Builder
	sb.Grow(len(lower))
	for _, r := range lower {
		if isVowel(r) {
			sb.WriteRune('*')
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// @leet end

// Keynold
