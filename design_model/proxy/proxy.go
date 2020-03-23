package proxy

import (
	"fmt"
	"strings"
)

/*
代理模式
1. 定义：
	给某一个对象提供一个代理，并由代理对象控制对原对象的引用。
2. 实现步骤
	定义接口
	定义实现对象
	定义代理实现对象
3. 使用场景
	并由代理对象控制对原对象的引用，增加请求注入劫持
4. 优点
	代理模式能够协调调用者和被调用者，在一定程度上降低了系统的耦合度。
5. 参考文章: https://studygolang.com/articles/7193
 */

// 在例子中,Github是被代理的对象,并实现了Git接口,每次调用都要经过GitBash,也就是可以在GitBash对访问进行控制,而当我们需要通过Gitlab来下载时,只要Gitlab实现Git接口,然后在GetGit替换Github即可.
type Git interface {
	Clone(url string) bool
}

type Github struct {}

func (*Github) Clone(url string) bool {
	if strings.HasPrefix(url,"https") {
		fmt.Println("clone from", url)
		return true
	}
	fmt.Println("failed to clone from", url)
	return false
}

type Coder struct {}

func (c *Coder) GetCode(url string) {
	gitBash := GetGit(1)
	if gitBash.Clone(url) {
		fmt.Println("success")
	}else {
		fmt.Println("failed")
	}
}

func GetGit(t int) Git {
	if t == 1 {
		return &GitBash{GitCmd: &Github{}}
	}
	return nil // may add gitlab
}

type GitBash struct {
	GitCmd Git
}

func (g *GitBash) Clone(url string) bool {
	return g.GitCmd.Clone(url)
}