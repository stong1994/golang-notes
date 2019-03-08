package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var urlKey = struct{}{}

// 利用context和goroutine来实现爬取网站，并设置超时
// 缺点：为了练习使用ctx的value方法，每个网站访问了多次
// 会在diff-url-data 中展示，如果一个数据有多个数据源，同时请求多个数据源，当一个goroutine获取到返回值时，终止其它goroutine
func main() {
	values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
	wg := &sync.WaitGroup{} // 如果wg需要作为参数传入函数，必须要用指针类型传递
	ctx, cancel := context.WithCancel(context.Background())
	for _, v := range values {
		wg.Add(1)
		sub := context.WithValue(ctx, urlKey, v)
		go getContext(sub, wg)
	}

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
		println("canceling...")
	}()

	wg.Wait()
}

func getContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			println("getContext received done signal")
			return
		default:
			wg.Add(1)
			url := ctx.Value(urlKey).(string)
			resp, err := http.Get(url)
			if err != nil || resp.StatusCode != http.StatusOK {
				println(err)
				return
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				println(err)
				return
			}
			sub := context.WithValue(ctx, urlKey, fmt.Sprintf("%s:%x", url, md5.Sum(bytes)))
			go printContent(sub, wg)
			resp.Body.Close()
			time.Sleep(2 * time.Second) // 因为在for循环中的select的default分支，所以需要sleep，不然会一直运行
		}
	}
}

func printContent(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			println("print receive done signal")
			return
		default:
			content := ctx.Value(urlKey).(string)
			println("printing...")
			println(string(content[:50]))
			time.Sleep(time.Second)
		}
	}
}
