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
	for i := 0; i < int(limitCount); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vo.Update(ip)
		}()
	}
	wg.Wait()
	if !vo.CheckIp(ip) {
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
