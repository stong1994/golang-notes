# 获取随机数
获取随机数一般有两种方式：
1. math.rand 
2. crypto.rand

> math.rand为伪随机数，即结果是确定的（通过一系列运算），因此需要每次更新种子，来保证“随机性”

> crypto.rand为真随机数

通过`benchcmp`来进行比较性能

生成过程：
1. `RandNum`函数中调用`MathRandNum`函数，然后执行命令`go test -bench="BenchmarkRandNum"  > math`
2. `RandNum`函数中调用`CryptoRandNum`函数，然后执行命令`go test -bench="BenchmarkRandNum"  > crypto`
3. 通过生成的基准测试文件进行比较，执行命令`benchcmp math crypto`

```
benchmark               old ns/op     new ns/op     delta
BenchmarkRandNum-16     17244         1179          -93.16%

benchmark               old allocs     new allocs     delta
BenchmarkRandNum-16     0              4              +Inf%

benchmark               old bytes     new bytes     delta
BenchmarkRandNum-16     0             56            +Inf%

```
其中`old`为`math`，`new`为`crypto`，可以看出，生成真随机数的时间要比伪随机数的时间快很多，竟然为93%。但是真随机数需要更多的内存分配。
