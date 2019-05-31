package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

// 官方例子
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	// 获取临时对象，没有的话会調用New()创建
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset() // 重置
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	// 将临时对象放回到 Pool 中
	bufPool.Put(b)
}

func main() {
	Log(os.Stdout, "path", "/search?q=flowers")
}
