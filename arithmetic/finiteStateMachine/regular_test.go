package main

import "testing"

func TestRegular(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			"match0",
			"abc",
			"a.c",
			true,
		},
		{
			"match1",
			"aa",
			"a",
			false,
		}, {
			"match2",
			"aa",
			"a*",
			true,
		}, {
			"match3",
			"ab",
			".*",
			true,
		},
		{
			"match4",
			"aab",
			"c*a*b",
			true,
		},
		{
			"match5",
			"mississippi",
			"mis*is*p*.",
			false,
		},
		{
			"match6",
			"",
			"c*",
			true,
		},
		{
			"math6.5",
			"baabbbaccbccacacc",
			"c*..b*a*a.*a..*c",
			true,
		},
		{
			"match7",
			"mississippi",
			"mis*is*ip*.",
			true,
		},
		{
			"match8",
			"aaa",
			"a*a",
			true,
		},
		{
			"match9",
			"aaa",
			"ab*ac*a",
			true,
		},
		{
			"match10",
			"aaa",
			"ab*a*c*a",
			true,
		},
		{
			"match11",
			"a",
			"ab*",
			true,
		},
		{
			"match12",
			"bbbba",
			".*a*a",
			true,
		},
		{
			"match13",
			"ab",
			".*..",
			true,
		},
		{
			"match14",
			"a",
			".*..a*",
			false,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := isMatch(v.s, v.p)
			if got != v.want {
				t.Fatalf("want %v but got %v", v.want, got)
			}
		})
	}
}
