package main

import "testing"

func TestHasTwoPairsOfLetters(t *testing.T) {

	type testcase struct {
		Name           string
		Entry          entryTwo
		ExpectedResult bool
	}

	testcases := []testcase{
		testcase{
			Name: "xyxy",
			Entry: entryTwo{
				Line: "xyxy",
			},
			ExpectedResult: true,
		},
		testcase{
			Name: "aabcdefgaa",
			Entry: entryTwo{
				Line: "aabcdefgaa",
			},
			ExpectedResult: true,
		},
		testcase{
			Name: "aaa",
			Entry: entryTwo{
				Line: "aaa",
			},
			ExpectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			if result := tc.Entry.HasTwoPairsOfLetters(); result != tc.ExpectedResult {
				t.Fatalf("expected %s to return %t on HasTwoPairsOfLetters, got: %t", tc.Entry.Line, tc.ExpectedResult, result)
			}
		})
	}
}

func TestHasRepeatingWithComboBreaker(t *testing.T) {

	type testcase struct {
		Name           string
		Entry          entryTwo
		ExpectedResult bool
	}

	testcases := []testcase{
		testcase{
			Name: "xyx",
			Entry: entryTwo{
				Line: "xyx",
			},
			ExpectedResult: true,
		},
		testcase{
			Name: "abcdefeghi",
			Entry: entryTwo{
				Line: "abcdefeghi",
			},
			ExpectedResult: true,
		},
		testcase{
			Name: "aaa",
			Entry: entryTwo{
				Line: "aaa",
			},
			ExpectedResult: true,
		},
		testcase{
			Name: "abc",
			Entry: entryTwo{
				Line: "abc",
			},
			ExpectedResult: false,
		},
		testcase{
			Name: "fawpewfmeaf",
			Entry: entryTwo{
				Line: "fawpewfmeaf",
			},
			ExpectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			if result := tc.Entry.HasRepeatingWithComboBreaker(); result != tc.ExpectedResult {
				t.Fatalf("expected %s to return %t on HasRepeatingWithComboBreaker, got: %t", tc.Entry.Line, tc.ExpectedResult, result)
			}
		})
	}
}
