package join_string

import (
	"bytes"
	"fmt"
	"strings"
)

func StringAdd(strs []string) string {
	var str string
	for i := range strs {
		str += strs[i]
	}
	return str
}

func SprintfAdd(a, b, c, d, e, f, g, h, i, j string) string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", a, b, c, d, e, f, g, h, i, j )
}

func StringJoin(strs []string) string {
	return 	strings.Join(strs, "")
}

func StringJoinWithAppend(strs []string) string {
	var arr = make([]string, 0, 10000)
	for i := range strs {
		arr = append(arr, strs[i])
	}
	return 	strings.Join(arr, "")
}

func BufferWith(strs []string) string {
	var buf bytes.Buffer
	for i := range strs {
		buf.WriteString(strs[i])
	}
	return buf.String()
}

func BufferWithGrow(strs []string) string {
	var buf bytes.Buffer
	buf.Grow(10000)
	for i := range strs {
		buf.WriteString(strs[i])
	}
	return buf.String()
}