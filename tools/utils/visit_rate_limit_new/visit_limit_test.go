package visit_limit

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


const (
	limitCount = 100
	limitDuration = 10 * 1000
	nodeCount = 10 // 节点数量
)

// 测试环形链表,每个节点不能为空
func TestNodeListValid(t *testing.T) {
	list := NewNodeList(nodeCount)
	node := list.headNode
	head := node
	if node == nil {
		t.Fatal("node is nil")
	}
	for i := 0; i < nodeCount*2; i++ {
		node = node.next
		if node == head {
			break
		}
		if node == nil {
			t.Fatal("node is nil")
		}
	}
}

// 测试环形链表，收尾节点是否一致
func TestNodeConsistency(t *testing.T) {
	list := NewNodeList(nodeCount)
	node := list.headNode
	first := node
	for i := 0; i < nodeCount; i++ {
		node = node.next
	}
	if first != node {
		t.Fatal("begin node and end node should be equal")
	}
}

// 测试同一个ip在10s内访问次数超过限制
func TestOneIpVisit1(t *testing.T) {
	ip := "11111"
	wg := sync.WaitGroup{}

	nl := NewVisitLimit(limitCount, limitDuration, nodeCount)
	for i := 0; i < limitCount+1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nl.UpdateIp(ip)
		}()
	}
	wg.Wait()
	if nl.CheckIP(ip) {
		t.Log("sum is ", nl.VisitSum(ip))
		t.Fatal("should be added in blacklist")
	} else {
		t.Log("success")
	}
}

// 测试同一个ip在10s内访问次数不超过限制
func TestOneIpVisit2(t *testing.T) {
	ip := "22222"
	wg := sync.WaitGroup{}
	nl := NewVisitLimit(limitCount, limitDuration, nodeCount)
	for i := 0; i < limitCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nl.UpdateIp(ip)
		}()
	}
	wg.Wait()
	if nl.CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", nl.VisitSum(ip))
		t.Fatal("should not be added in blacklist")
	}
}

// 测试同一个ip在10s内访问次数超过限制后是否在黑名单
func TestShouldInBlankList(t *testing.T) {
	ip := "11111"
	wg := sync.WaitGroup{}
	nl := NewVisitLimit(limitCount, limitDuration, nodeCount)
	for i := 0; i < limitCount+1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nl.UpdateIp(ip)
		}()
	}
	wg.Wait()
	if !nl.CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", nl.VisitSum(ip))
		t.Fatal("should be added in blacklist")
	}
}

// 测试同一个ip在10s内访问次数不超过限制后是否在黑名单
func TestShouldNotInBlankList(t *testing.T) {
	ip := "22222"
	wg := sync.WaitGroup{}
	nl := NewVisitLimit(limitCount, limitDuration, nodeCount)
	for i := 0; i < limitCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nl.UpdateIp(ip)
		}()
	}
	wg.Wait()
	if nl.CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", nl.VisitSum(ip))
		t.Fatal("should not be added in blacklist")
	}
}

// 测试同一个ip在10s内访问次数不超过限制后是否在黑名单
func TestVisit101Dur11s(t *testing.T) {
	ip := "22222"
	wg := sync.WaitGroup{}
	nl := NewVisitLimit(limitCount, limitDuration, nodeCount)
	for i := 0; i < limitCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nl.UpdateIp(ip)
		}()
	}
	wg.Wait()
	time.Sleep(10 * time.Second)
	nl.UpdateIp(ip)

	if nl.CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", nl.VisitSum(ip))
		t.Fatal("should not be added in blacklist")
	}
}

// 测试同一个ip在1s访问10次是否在黑名单
func TestVisit10Dur1s(t *testing.T) {
	ip := "22222"
	nl := NewVisitLimit(limitCount, limitDuration, nodeCount)
	for i := 0; i < nodeCount+2; i++ {
		wg := sync.WaitGroup{}
		for i := 0; i < nodeCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				nl.UpdateIp(ip)
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

	if nl.CheckIP(ip) {
		t.Log("success")
	} else {
		t.Log("sum is ", nl.VisitSum(ip))
		t.Fatal("should not be added in blacklist")
	}
}

// 测试同一个ip在1s访问12次是否在黑名单
func TestVisit11Dur1s(t *testing.T) {
	ip := "22222"
	nl := NewVisitLimit(limitCount, limitDuration, nodeCount)
	for i := 0; i < nodeCount; i++ {
		wg := sync.WaitGroup{}
		for i := 0; i < nodeCount+2; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				nl.UpdateIp(ip)
			}()
		}
		wg.Wait()
		time.Sleep(time.Second)
	}
	t.Log("sum is", nl.VisitSum(ip))
	if !nl.CheckIP(ip) {
		t.Log("success")
	} else {
		t.Fatal("should be added into blacklist")
	}
}
//
//
//// 测试同一个ip在10s后访问的节点分布情况
//func TestVisitAfter20s(t *testing.T) {
//	ip := "22222"
//	Update(ip)
//	time.Sleep(20 * time.Second)
//	for i := 0; i < 11; i++ {
//		wg := sync.WaitGroup{}
//		for i := 0; i < nodeCount; i++ {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				Update(ip)
//			}()
//		}
//		wg.Wait()
//		time.Sleep(time.Second)
//		for i := 10; i > 0; i-- {
//			ipLimitInfos.locker.RLock()
//			fmt.Printf("%d \t", ipLimitInfos.ips[ip].headNode.num)
//			ipLimitInfos.ips[ip].headNode = ipLimitInfos.ips[ip].headNode.next
//			ipLimitInfos.locker.RUnlock()
//		}
//		fmt.Println()
//	}
//	if CheckIP(ip) {
//		t.Log("success")
//	} else {
//		t.Log("sum is", Sum(ip))
//		t.Fatal("should be added into blacklist")
//	}
//}

// 测试第一秒访问10次，然后第2-10秒访问90次,然后第11秒访问1次，也就是在最新的10秒内访问了91次，查看访问次数是否为91
func TestValidAfter10Sv1(t *testing.T) {
	ip := "11111"
	nl := NewVisitLimit(limitCount, limitDuration, nodeCount)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nl.UpdateIp(ip)
		}()
	}
	time.Sleep(time.Second)
	wg.Wait()

	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nl.UpdateIp(ip)
		}()
	}
	time.Sleep(time.Second * 9)
	wg.Wait()

	nl.UpdateIp(ip)

	if nl.CheckIP(ip) && nl.VisitSum(ip) == 91 {
		t.Log("success")
	} else {
		t.Log(nl.VisitSum(ip))
		t.Fatal("should be valid")
	}
}

func printNodes(nodes *visitLimit, ip string)  {
	node := nodes.ipLimit.ips[ip].headNode
	for i := 0; i < nodeCount; i++ {
		fmt.Println(node.num)
		node = node.next
	}
}