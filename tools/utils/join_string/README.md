# 比较字符串拼接速度

#### 1. 拼接字符串的几种方法：
1. 直接拼接
2. 利用fmt.Sprintf()
3. strings.Join()
3. buffer.WriteString()

> 数据量小的话，其实都无所谓，如果数据量大的话，fmt.Sprintf()肯定不方便。剩下的
三个里面，直接拼接性能最差，如果是现成的数组，用strings.Join()，否则用buffer.WriteString()

#### 2. byte数组转字符串进行拼接
```go
var s string
var x = []byte{1023:'x'}
var y = []byte{1023:'y'}
s = (" " + string(x) + string(y))[1:]
```
> 这样不会有底层数组的复制,比直接进行拼接更高效