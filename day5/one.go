package main

import "strings"

type entryOne struct {
	Line string
}

func (e entryOne) Nice() bool {
	return e.HasAtLeast3Vowels() && e.HasLetterTwiceInARow() && !e.HasNaughtyCombo()
}

func (e entryOne) HasAtLeast3Vowels() bool {
	const vowels = "aeiou"
	var count int
	for _, l := range e.Line {
		for _, vowel := range vowels {
			if l == vowel {
				count++
			}
		}
	}
	return count >= 3
}

func (e entryOne) HasLetterTwiceInARow() bool {
	var prev rune
	for _, l := range e.Line {
		if l == prev {
			return true
		}
		prev = l
	}
	return false
}

func (e entryOne) HasNaughtyCombo() bool {
	for _, naughty := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(e.Line, naughty) {
			return true
		}
	}
	return false
}
