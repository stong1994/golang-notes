package parallel_access

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	urlChan      = make(chan string)
	dataData     = make(chan []byte)
	errChan      = make(chan error)
	timeoutError = errors.New("time out")
)

/*
用于请求多个数据源，当获取到第一个返回值后，取消其它goroutine，节省资源。
*/
func DispatcherUrl(timeOut time.Duration, urls ...string) ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for _, url := range urls {
			sub, _ := context.WithTimeout(ctx, timeOut) // 不用获取cancel函数，根ctx的额cancel函数的调用会触发子ctx的cancel函数
			go GetData(sub)                             // 一个url 对应一个goroutine
			urlChan <- url
		}
	}()
	select {
	case data := <-dataData:
		// 停止其它数据来源的获取
		cancel()
		return data, nil
	case <-time.After(timeOut): // 超时限制
		cancel()
		return nil, timeoutError
	}
}

func GetData(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case url := <-urlChan:
			resp, err := http.Get(url)
			if err != nil {
				errChan <- err
				return
			}
			if resp.StatusCode != http.StatusOK {
				errChan <- errors.New(fmt.Sprintf("response code is %d", resp.StatusCode))
				return
			}
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				errChan <- err
				return
			}
			dataData <- data
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
