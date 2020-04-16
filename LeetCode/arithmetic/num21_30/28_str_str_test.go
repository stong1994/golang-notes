package num21_30

import "testing"

/*
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
*/
func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{
			"test1",
			"hello",
			"ll",
			2,
		},
		{
			"test2",
			"aaaaa",
			"bba",
			-1,
		},
		{
			"test3",
			"mississippi",
			"pi",
			9,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := strStr(v.haystack, v.needle)
			if got != v.want {
				t.Fatalf("want %d got %d", v.want, got)
			}
		})
	}
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] {
			for j := range needle {
				if needle[j] != haystack[i+j] {
					goto Continue
				}
			}
			return i
		}
	Continue:
	}
	return -1
}
