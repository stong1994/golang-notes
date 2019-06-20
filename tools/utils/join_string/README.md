# 比较字符串拼接速度

拼接字符串的几种方法：
1. 直接拼接
2. 利用fmt.Sprintf()
3. strings.Join()
3. buffer.WriteString()

> 数据量小的话，其实都无所谓，如果数据量大的话，fmt.Sprintf()肯定不方便。剩下的
三个里面，直接拼接性能最差，如果是现成的数组，用strings.Join()，否则用buffer.WriteString()