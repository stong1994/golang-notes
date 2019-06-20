# 获取随机数
获取随机数一般有两种方式：
1. math.rand 
2. crypto.rand

> math.rand为伪随机数，即结果是确定的（通过一系列运算），因此需要每次更新种子，来保证“随机性”

> crypto.rand为真随机数

经过性能测试，crypto.math生成随机数的速度竟然比math.rand生成随机数的速度快，但是要消耗的内存更多。