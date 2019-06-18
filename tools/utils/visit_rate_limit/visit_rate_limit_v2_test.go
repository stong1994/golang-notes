package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const SUCCESS = "success"

// 测试10s内访问101次
func TestAccess101Dur10S(t *testing.T) {
	vo := NewVisitOperation()
	vo.Start()

	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < int(limitCount)+1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.Update(ip)
		}()
	}
	wg.Wait()
	if vo.CheckIp(ip) {
		t.Log(vo.Sum(ip))
		t.Fatal("can not access 101 times duration 10s")
	} else {
		t.Log(SUCCESS)
	}
}

// 测试10s内访问100次
func TestAccess100Dur10S(t *testing.T) {
	vo := NewVisitOperation()
	vo.Start()

	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < limitCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.Update(ip)
		}()
	}
	wg.Wait()
	if !vo.CheckIp(ip) && vo.Sum(ip) != limitCount {
		t.Log(vo.Sum(ip))
		t.Fatal("should be valid")
	} else {
		t.Log(SUCCESS)
	}
}

// 测试每秒访问10次
func TestAccess10PerSecond(t *testing.T) {
	vo := NewVisitOperation()
	vo.Start()

	ip := "11111"
	for i := 0; i < 11; i++ {
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				vo.Update(ip)
			}()
		}
		time.Sleep(time.Second)
		wg.Wait()
		fmt.Println(vo.Sum(ip))
	}
	if vo.CheckIp(ip) {
		t.Log(SUCCESS)
	} else {
		t.Log(vo.Sum(ip))
		t.Fatal("should be valid")
	}
}

// 测试每秒访问11次
func TestAccess11PerSecond(t *testing.T) {
	vo := NewVisitOperation()
	vo.Start()

	ip := "11111"
	for i := 0; i < 11; i++ {
		wg := sync.WaitGroup{}
		for i := 0; i < 11; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				vo.Update(ip)
			}()
		}
		time.Sleep(time.Second)
		wg.Wait()
		fmt.Println(vo.Sum(ip))
	}
	if !vo.CheckIp(ip) {
		t.Log(SUCCESS)
	} else {
		t.Log(vo.Sum(ip))
		t.Fatal("should not be valid")
	}
}

// 测试第一秒访问100次，然后第11秒访问1次
func TestValidAfter10S(t *testing.T) {
	vo := NewVisitOperation()
	vo.Start()

	ip := "11111"

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.Update(ip)
		}()
	}
	wg.Wait()
	fmt.Println(vo.Sum(ip))

	time.Sleep(time.Second * 10)

	vo.Update(ip)

	if vo.CheckIp(ip) && vo.Sum(ip) == 1 {
		t.Log(SUCCESS)
	} else {
		t.Log(vo.Sum(ip))
		t.Fatal("should be valid")
	}
}

// 测试第一秒访问10次，然后第2秒访问90次,然后第11秒访问1次
func TestValidAfter10Sv2(t *testing.T) {
	vo := NewVisitOperation()
	vo.Start()

	ip := "11111"

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.Update(ip)
		}()
	}
	time.Sleep(time.Second)
	wg.Wait()

	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.Update(ip)
		}()
	}
	time.Sleep(time.Second * 9)
	wg.Wait()

	vo.Update(ip)

	if vo.CheckIp(ip) && vo.Sum(ip) == 91 {
		t.Log(SUCCESS)
	} else {
		t.Log(vo.Sum(ip))
		t.Fatal("should be valid")
	}
}

// 测试第一秒访问100次，然后第21秒访问1次
func TestValidAfter20S(t *testing.T) {
	vo := NewVisitOperation()
	vo.Start()

	ip := "11111"

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.Update(ip)
		}()
	}
	wg.Wait()
	fmt.Println(vo.Sum(ip))

	time.Sleep(time.Second * 20)

	vo.Update(ip)

	if vo.CheckIp(ip) && vo.Sum(ip) == 1 {
		t.Log(SUCCESS)
	} else {
		t.Log(vo.Sum(ip))
		t.Fatal("should be valid")
	}
}

// 性能测试 v1-读 /
// BenchmarkV1Read-16    	2000000000	          0.03 ns/op
func BenchmarkV1Read(b *testing.B) {
	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CheckIP(ip)
		}()
	}
	wg.Wait()
}

// 性能测试 v2-读 /
// BenchmarkV2Read-16    	2000000000	         0.09 ns/op
func BenchmarkV2Read(b *testing.B) {
	vo := NewVisitOperation()
	vo.Start()

	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.CheckIp(ip)
		}()
	}
	wg.Wait()
}

// 性能测试 v1-写 /
// BenchmarkV1Write-16    	2000000000	        0.03 ns/op
func BenchmarkV1Write(b *testing.B) {
	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Update(ip)
		}()
	}
	wg.Wait()
}

// 性能测试 v2-写 /
// BenchmarkV2Write-16    	2000000000	         0.10 ns/op
func BenchmarkV2Write(b *testing.B) {
	vo := NewVisitOperation()
	vo.Start()

	b.ResetTimer()

	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.Update(ip)
		}()
	}
	wg.Wait()
}

// 性能测试 v1-读写 /
// BenchmarkV1RW-16    	       2	 656289300 ns/op	32007392 B/op	  166990 allocs/op
func BenchmarkV1RW(b *testing.B) {
	b.ReportAllocs()
	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < 500000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CheckIP(ip)
		}()
	}
	for i := 0; i < 500000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Update(ip)
		}()
	}

	wg.Wait()
}

// 性能测试 v2-读写 /
// BenchmarkV2RW-16    	       1	1952373500 ns/op	164446384 B/op	 1938529 allocs/op
func BenchmarkV2RW(b *testing.B) {
	vo := NewVisitOperation()
	vo.Start()
	b.ReportAllocs()
	b.ResetTimer()

	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < 500000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.CheckIp(ip)
		}()
	}
	for i := 0; i < 500000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.Update(ip)
		}()
	}
	wg.Wait()
}