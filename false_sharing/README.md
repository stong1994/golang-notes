# false sharing 

*推荐文章：[false sharing 与 CPU缓存设计](https://www.yuque.com/docs/share/939f249c-255a-4ba7-a58c-3ffb6eb7f782)*

### 记录要点：
1. 空间局部性理论：当 CPU 访问一个变量时，它可能很快就会读取它旁边的变量
2. CPU 中缓存的最小化单位是缓存行
3. 缓存 false sharing：一个 CPU 核更新变量会强制其他 CPU 核更新缓存
4. false sharing带来一个问题：假设core1，core2中的缓存行有a,b两个变量，core1中的a更新。即使b没有更新，当core2再次访问b时,缓存失效，
需要再次加载缓存，带来性能问题。
5. 解决上述问题的常用方法是**缓存填充**：在变量之间填充一些无意义的变量。使一个变量单独占用 CPU 核的缓存行，因此当其他核更新时，其他变量不会使该核从内存中重新加载变量。

### 实验结果
文章中给出的填充缓存结构为
```go
type Pad struct {
	a   uint64
	_p1 [8]uint64
	b   uint64
	_p2 [8]uint64
	c   uint64
	_p3 [8]uint64
}
```
注意这里填充的是8个8字节数据，加起来正好是64字节，即CPU的一个缓存行大小。  
测试结果：
```
BenchmarkPad-16         2000000000               0.02 ns/op  
```
我将填充数据改为7个8字节数据，即
```go
type Pad struct {
	a   uint64
	_p1 [7]uint64
	b   uint64
	_p2 [7]uint64
	c   uint64
	_p3 [7]uint64
}
```
测试结果与上个测试相同。  
**分析原因**：第一个测试填充数据为64字节，那么导致每个数据单独占用一个缓存行（因此第二个数据放不下了）；第二个测试填充字节为56，实际用到的数据大小为8字节，正好是一个缓存行。每个实际用到的数据都存在不同的缓存行，解决了`false sharing`带来的问题。因此这两个测试结果相同。

为了证明上述原因，我又把填充数据改为了6个8字节数据，即
```go
type Pad struct {
	a   uint64
	_p1 [6]uint64
	b   uint64
	_p2 [6]uint64
	c   uint64
	_p3 [6]uint64
}
```
分析可知，a和b将会在同一个缓存行。测试结果的性能应该小于上述两个测试。
```
BenchmarkPad-16         2000000000               0.03 ns/op
```
果然性能下降！

再将填充数据改为5个8字节数据，即
```go
type Pad struct {
	a   uint64
	_p1 [5]uint64
	b   uint64
	_p2 [5]uint64
	c   uint64
	_p3 [5]uint64
}
```
a和b还是在同一个缓存行，测试结果应和上一个测试结果相同
```
BenchmarkPad-16         2000000000               0.03 ns/op 
```
完美！！

