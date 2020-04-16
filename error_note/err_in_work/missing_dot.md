报错信息：
```
malformed module path "awesome-dot/sub_util": missing dot in first path element
```
##### 第一种可能：项目结构

```
about_dot
    sub_server
        main.go
    sub_util
        util.go
```
其中，`main.go`中引用了`util.go`中的函数。
在`sub_server`目录下执行
```
go mod init xxx
go mod tidy
```
会报此错。 

按照错误提示，将`about_dot`中增加`.`,即改为`about_dot.local`，那么重复上述步骤，会报错
``` 
module about-dot.local/sub_util: Get https://proxy.golang.org/about-dot.local/sub_util/@v/list: dial tcp 172.217.24.17:443: i/o timeout
```
即发生了错误的远程调用。

正确做法：

> 那么在这种情况下，直接在项目最外层执行`go mod init`即可 



##### 第二种可能：GO111MODULE设置为了on

解决办法：` export GO111MODULE=off`

据某位大佬所说，这种错误只会在gomodule开启才会产生，所以关闭了也就不报错了