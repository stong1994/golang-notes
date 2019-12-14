package num21_30

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

/*
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
*/
func TestFindSubstring(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  []int
	}{
		{
			"test1",
			"barfoothefoobarman",
			[]string{"foo", "bar"},
			[]int{0, 9},
		}, {
			"test2",
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "word"},
			[]int{},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := findSubstring(v.s, v.words)
			if !assert.IsEqual(got, v.want) {
				t.Fatalf("want %v got %v", v.want, got)
			}
		})
	}
}

// 暴力破解法
/*func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}

	result := make([]int, 0)
	wordLen := len(words[0])
	for i := 0; i < len(s) - wordLen + 1; i++ {
		for j := 0; j < len(words); j++ {
			if s[i:i+wordLen] == words[j] { // 首次匹配字符串
				//wordsCopy := append(words[:j], words[j+1:]...)
				wordsCopy := make([]string, len(words))
				copy(wordsCopy, words)
				wordsCopy = append(wordsCopy[:j], wordsCopy[j+1:]...)
				//_ = wordsCopy
				if fixStr(s[i+wordLen:], wordsCopy, wordLen) {
					result = append(result, i)
				}
				break
			}
		}
	}
	return result
}

func fixStr(s string, words []string, wordLen int) bool {
	if len(words) == 0 {
		return true
	}
	if len(s) < wordLen {
		return false
	}
	for i := 0; i < len(words); i++ {
		if s[:wordLen] == words[i] {
			words = append(words[:i], words[i+1:]...)
			return fixStr(s[wordLen:], words, wordLen)
		}
	}
	return false
}*/

// hash map
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}

	wordsMap := make(map[string]int)
	for _, v := range words {
		wordsMap[v] = wordsMap[v] + 1
	}

	result := make([]int, 0)
	wordLen := len(words[0])
	totalLen := len(words) * wordLen
	for i := 0; i < len(s)-wordLen+1; i++ {
		for j := 0; j < len(words); j++ {
			if s[i:i+wordLen] == words[j] { // 首次匹配字符串
				wordsCopy := make(map[string]int, len(wordsMap))
				for k, v := range wordsMap {
					wordsCopy[k] = v
				}
				if fixStr(wordsCopy, s[i:], wordLen, totalLen) {
					result = append(result, i)
				}
				break
			}
		}
	}
	return result
}

func fixStr(wordsMap map[string]int, s string, wordLen int, totalLen int) bool {
	if len(s) < totalLen {
		return false
	}
	for i := 0; i < totalLen/wordLen; i++ {
		key := s[i*wordLen : (i+1)*wordLen]
		if val, ok := wordsMap[key]; ok && val > 0 {
			wordsMap[key] = val - 1
		} else {
			return false
		}
	}
	return true
}
