package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 测试环形链表,每个节点不能为空
func TestNodeListValid(t *testing.T) {
	list := NewNodeList(time.Now().UnixNano() / 1e6)
	node := list.headNode
	if node == nil {
		t.Fatal("node is nil")
	}
	for i := 0; i < nodeCount*2; i++ {
		node = node.next
		if node == nil {
			t.Fatal("node is nil")
		}
	}
}

// 测试环形链表，收尾节点是否一致
func TestNodeConsistency(t *testing.T) {
	list := NewNodeList(time.Now().UnixNano() / 1e6)
	node := list.headNode
	first := node
	for i := 0; i < nodeCount; i++ {
		node = node.next
	}
	if first != node {
		t.Fatal("began and end should equal")
	}
}

// 测试同一个ip在10s内访问次数超过限制
func TestOneIpVisit1(t *testing.T) {
	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < limitCount+1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Update(ip)
		}()
	}
	wg.Wait()
	if CheckIP(ip) {
		t.Log("sum is ", Sum(ip))
		t.Fatal("should be added in blacklist")
	} else {
		t.Log("success")
	}
}

// 测试同一个ip在10s内访问次数不超过限制
func TestOneIpVisit2(t *testing.T) {
	ip := "22222"
	wg := sync.WaitGroup{}
	for i := 0; i < limitCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Update(ip)
		}()
	}
	wg.Wait()
	if CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", Sum(ip))
		t.Fatal("should not be added in blacklist")
	}
}

// 测试同一个ip在10s内访问次数超过限制后是否在黑名单
func TestShouldInBlankList(t *testing.T) {
	ip := "11111"
	wg := sync.WaitGroup{}
	for i := 0; i < limitCount+1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Update(ip)
		}()
	}
	wg.Wait()
	CheckIP(ip)
	if !inBlankList(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", Sum(ip))
		t.Fatal("should be added in blacklist")
	}
}

// 测试同一个ip在10s内访问次数不超过限制后是否在黑名单
func TestShouldNotInBlankList(t *testing.T) {
	ip := "22222"
	wg := sync.WaitGroup{}
	for i := 0; i < limitCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Update(ip)
		}()
	}
	wg.Wait()
	CheckIP(ip)
	if !inBlankList(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", Sum(ip))
		t.Fatal("should not be added in blacklist")
	}
}

// 测试同一个ip在10s内访问次数不超过限制后是否在黑名单
func TestVisit101Dur11s(t *testing.T) {
	ip := "22222"
	wg := sync.WaitGroup{}
	for i := 0; i < limitCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Update(ip)
		}()
	}
	wg.Wait()
	time.Sleep(10 * time.Second)
	Update(ip)

	if CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", Sum(ip))
		t.Fatal("should not be added in blacklist")
	}
}

// 测试同一个ip在1s访问10次是否在黑名单
func TestVisit10Dur1s(t *testing.T) {
	ip := "22222"
	for i := 0; i < nodeCount+2; i++ {
		wg := sync.WaitGroup{}
		for i := 0; i < nodeCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				Update(ip)
			}()
		}
		wg.Wait()
		time.Sleep(time.Second)
		/*for i := 10; i > 0; i-- {
			ipm.locker.RLock()
			fmt.Printf("%d \t", ipm.ips[ip].headNode.num)
			ipm.ips[ip].headNode = ipm.ips[ip].headNode.next
			ipm.locker.RUnlock()
		}
		fmt.Println()*/
	}

	if CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", Sum(ip))
		t.Fatal("should not be added in blacklist")
	}
}

// 测试同一个ip在1s访问12次是否在黑名单
func TestVisit11Dur1s(t *testing.T) {
	ip := "22222"
	for i := 0; i < nodeCount+4; i++ {
		wg := sync.WaitGroup{}
		for i := 0; i < nodeCount+1; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				Update(ip)
			}()
		}
		wg.Wait()
		time.Sleep(time.Second)
		/*for i := 10; i > 0; i-- {
			ipm.locker.RLock()
			fmt.Printf("%d \t", ipm.ips[ip].headNode.num)
			ipm.ips[ip].headNode = ipm.ips[ip].headNode.next
			ipm.locker.RUnlock()
		}
		fmt.Println()*/
	}
	if !CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is", Sum(ip))
		t.Fatal("should be added into blacklist")
	}
}

// 测试同一个ip在10s后访问的节点分布情况
func TestVisitAfter10s(t *testing.T) {
	ip := "22222"
	Update(ip)
	time.Sleep(10 * time.Second)
	for i := 0; i < 11; i++ {
		wg := sync.WaitGroup{}
		for i := 0; i < nodeCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				Update(ip)
			}()
		}
		wg.Wait()
		time.Sleep(time.Second)
		for i := 10; i > 0; i-- {
			ipm.locker.RLock()
			fmt.Printf("%d \t", ipm.ips[ip].headNode.num)
			ipm.ips[ip].headNode = ipm.ips[ip].headNode.next
			ipm.locker.RUnlock()
		}
		fmt.Println()
	}
	if CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is", Sum(ip))
		t.Fatal("should be added into blacklist")
	}
}

// 测试同一个ip在10s后访问的节点分布情况
func TestVisitAfter20s(t *testing.T) {
	ip := "22222"
	Update(ip)
	time.Sleep(20 * time.Second)
	for i := 0; i < 11; i++ {
		wg := sync.WaitGroup{}
		for i := 0; i < nodeCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				Update(ip)
			}()
		}
		wg.Wait()
		time.Sleep(time.Second)
		for i := 10; i > 0; i-- {
			ipm.locker.RLock()
			fmt.Printf("%d \t", ipm.ips[ip].headNode.num)
			ipm.ips[ip].headNode = ipm.ips[ip].headNode.next
			ipm.locker.RUnlock()
		}
		fmt.Println()
	}
	if CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is", Sum(ip))
		t.Fatal("should be added into blacklist")
	}
}
