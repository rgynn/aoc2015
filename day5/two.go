package main

import (
	"fmt"
	"strings"
)

type entryTwo struct {
	Line string
}

func (e entryTwo) Nice() bool {
	return e.HasTwoPairsOfLetters() && e.HasRepeatingWithComboBreaker()
}

// It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
func (e entryTwo) HasTwoPairsOfLetters() bool {

	for i := range e.Line {
		if len(e.Line) == i+1 {
			break
		}
		pair := fmt.Sprintf("%s%s", string(e.Line[i]), string(e.Line[i+1]))
		if count := strings.Count(e.Line, pair); count > 1 {
			return true
		}
	}

	return false
}

// It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
func (e entryTwo) HasRepeatingWithComboBreaker() bool {

	if len(e.Line) == 3 {
		if e.Line[0] == e.Line[2] {
			return true
		}
	}

	if len(e.Line) > 3 {
		for i := range e.Line {
			if len(e.Line) == i+2 {
				break
			}
			if e.Line[i] == e.Line[i+2] {
				return true
			}
		}
	}

	return false
}
