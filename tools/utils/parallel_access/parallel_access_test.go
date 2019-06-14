package parallel_access

import (
	"testing"
	"time"
)

func TestSameDataWithDiffUrl(t *testing.T) {
	addrs := []string{"https://www.baidu.com", "https://www.zhihu.com"}
	data, err := DispatcherUrl(time.Second, addrs...)
	if err != nil {
		t.Error("get error", err)
		return
	}
	if len(data) > 0 {
		t.Log("pass")
	} else {
		t.Error("error: len of data is 0")
	}
}

func TestSameDataWithTimeout(t *testing.T) {
	addrs := []string{"https://www.baidu.com", "https://www.zhihu.com"}
	_, err := DispatcherUrl(time.Microsecond, addrs...)
	if err != nil {
		t.Log("pass")
		return
	}
	t.Error("expect timeout but not")
}

func Benchmark_SameDataWithDiffUrl(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	addrs := []string{"https://www.baidu.com", "https://www.zhihu.com"}
	for i := 0; i < 100; i++ {
		data, err := DispatcherUrl(time.Second, addrs...)
		if err != nil {
			b.Fatal(err)
		}
		if len(data) == 0 {
			b.Fatal("len of data is 0")
		}
	}
}
