package main

import (
	"sort"
	"strconv"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{
			"test1",
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"test2",
			"2",
			[]string{"a", "b", "c"},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := letterCombinations(v.arg)
			if !judgeSameArr(got, v.want) {
				t.Errorf("want %v got %v", v.want, got)
			}
		})
	}
}

func judgeSameArr(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	for i, v := range arr1 {
		if arr2[i] != v {
			return false
		}
	}
	return true
}

func letterCombinations(digits string) []string {
	parseInt, err := strconv.ParseInt(digits, 10, 64)
	if err != nil {
		return nil
	}
	digistInts := []int64{}
	for {
		n := parseInt / 10
		digistInts = append(digistInts, parseInt-n*10)
		if n == 0 {
			break
		}
		parseInt = n
	}
	digistsMap := map[int64][]byte{2: []byte{'a', 'b', 'c'}, 3: []byte{'d', 'e', 'f'}, 4: []byte{'g', 'h', 'i'}, 5: []byte{'j', 'k', 'l'},
		6: []byte{'m', 'n', 'o'}, 7: []byte{'p', 'q', 'r', 's'}, 8: []byte{'t', 'u', 'v'}, 9: []byte{'w', 'x', 'y', 'z'}}

	var res []string
	var valid [][]byte
	for _, v := range digistInts {
		val, ok := digistsMap[v]
		if !ok {
			continue
		}
		valid = append(valid, val)
	}

	resByte := combineArr(valid)
	for _, v := range resByte {
		res = append(res, string(v))
	}
	return res
}

func combineArr(data [][]byte) [][]byte {
	l := len(data)
	var res [][]byte
	if l == 1 {
		for _, v := range data[0] {
			res = append(res, []byte{v})
		}
		return res
	}
	if l == 2 {
		for _, v := range data[1] {
			for _, b := range data[0] {
				res = append(res, []byte{v, b})
			}
		}
		return res
	}
	left := combineArr(data[0 : l/2])
	right := combineArr(data[l/2:])
	for _, v := range left {
		for _, r := range right {
			res = append(res, append(r, v...))
		}
	}
	return res
}
