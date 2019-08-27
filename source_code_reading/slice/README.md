# 分析slice源码

添加代码
```go
func appendSlice()  {
	var arr []int
	arr = append(arr, 1)
}
```
编译得
```
...
0x0042 00066 (append_slice.go:5)        CALL    runtime.growslice(SB)
''' 
```
找到`runtime/slice.go`.发现一共也就两百多行,一块看了得了.

#### 首先是结构体
```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
emmm,只有三个属性,对比`map`(`map`有九个属性),`slice`应该会简单很多.

#### 初始化slice
```go
func makeslice(et *_type, len, cap int) unsafe.Pointer {
	mem, overflow := math.MulUintptr(et.size, uintptr(cap))
	if overflow || mem > maxAlloc || len < 0 || len > cap {
		// NOTE: Produce a 'len out of range' error instead of a
		// 'cap out of range' error when someone does make([]T, bignumber).
		// 'cap out of range' is true too, but since the cap is only being
		// supplied implicitly, saying len is clearer.
		// See golang.org/issue/4085.
		mem, overflow := math.MulUintptr(et.size, uintptr(len))
		if overflow || mem > maxAlloc || len < 0 {
			panicmakeslicelen()
		}
		panicmakeslicecap()
	}

	return mallocgc(mem, et, true)
}
func panicmakeslicelen() {
	panic(errorString("makeslice: len out of range"))
}
func panicmakeslicecap() {
	panic(errorString("makeslice: cap out of range"))
}
```
1. 根据切片中存储的数据类型和容量`cap`来判断是否超出了内存限制,再根据长度`len`和数据类型来判断报`cap`超出范围还是`len`超出范围
2. 根据数据类型和`cap`给`slice`分配内存

我理所应当的在这段代码
```go
func appendSlice()  {
	var arr []int
	arr = append(arr, 1)
}
```
编译后的文件中查找这个函数`makeslice`,奇怪的是并没有找到,关于`arr`我只找到了这段
```
"".arr SBSS size=24
```
`size`为24字节很容易理解,因为我的电脑是64位,因此`len`和`cap`都占8字节,`unsafe.Pointer`为`unsafe.ArbitraryType`的别名,`unsafe.ArbitraryType`
又是`int`的别名,因此也是8字节,所以`arr`大小为24字节.

回到问题:为什么没有调用`makeslice`函数呢?

于是我用`make`初始化的方式来初始化`slice`
```go
var declare_make = make([]int, 5, 10)
```
编译得到
```
...
CALL    runtime.makeslice(SB)
...
CMPL    runtime.writeBarrier(SB), $0
...
CALL    runtime.gcWriteBarrier(SB)
...
```
不仅得到了预想中的`makeslice`函数,还得到了`写屏障`和`gc`的`写屏障`.不过今天只研究`slice`

处于闲的蛋疼的精神,我又写了一段初始化切片的代码
```go
var slice_init = []int{1}
```
编译结果
```
...
"".slice_init SDATA size=24
... 
```
和只声明类型的操作结果差不多,只不过一个是`SBSS`,一个是`SDATA`,即一个是未初始化,一个为已初始化.

所以得到**结论**:  
在对`slice`的操作中,只有用`make`进行初始化(编译得到的也是`SBSS`,即未初始化),才能调用`makeslice`函数.  

上述结论也很容易理解,因为只有在定义了`len`和`cap`的时候,才会涉及到超出内存限制的问题.

下面还有这段代码
```go
func makeslice64(et *_type, len64, cap64 int64) unsafe.Pointer {
	len := int(len64)
	if int64(len) != len64 {
		panicmakeslicelen()
	}

	cap := int(cap64)
	if int64(cap) != cap64 {
		panicmakeslicecap()
	}

	return makeslice(et, len, cap)
}
```
比上面的函数只是多了一个64,那么意思就是在指定`len`和`cap`的为`int64`类型时调用该函数?  
测试了一下.

代码
```go
var declare_make_64 = make([]int, int64(5), int64(10))
```
编译结果仍然调用的`makesilce`函数.
```
0x003a 00058 (declare_slice.go:16)      CALL    runtime.makeslice(SB)
```
算了,不管这个函数了.

#### 扩容
```go
func growslice(et *_type, old slice, cap int) slice {
	if raceenabled {
		callerpc := getcallerpc()
		racereadrangepc(old.array, uintptr(old.len*int(et.size)), callerpc, funcPC(growslice))
	}
	if msanenabled {
		msanread(old.array, uintptr(old.len*int(et.size)))
	}

	if cap < old.cap {
		panic(errorString("growslice: cap out of range"))
	}
    // todo 在什么情况下字段占用的大小为0? struct{}
	if et.size == 0 {
		// append should not create a slice with nil pointer but non-zero len.
		// We assume that append doesn't need to preserve old.array in this case.
		return slice{unsafe.Pointer(&zerobase), old.len, cap}
	}
    
	// 如果指定的cap大于两倍的slice的cap,则新的cap为指定的cap;否则如果slice的长度小于1024,那么新的cap为之前cap的两倍;如果不小于1024,
	// 则每次增加四分之一的cap,直到接近指定的cap
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.len < 1024 {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}

	// 计算出新的slice内存是否溢出,旧的len占用的内存大小,新的len占用的内存大小,cap占用的内存大小
	var overflow bool
	var lenmem, newlenmem, capmem uintptr
        // Specialize for common values of et.size.
        // For 1 we don't need any division/multiplication.
        // For sys.PtrSize, compiler will optimize division/multiplication into a shift by a constant.
        // For powers of 2, use a variable shift.
	switch {
	case et.size == 1:
		lenmem = uintptr(old.len)
		newlenmem = uintptr(cap)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem)
	case et.size == sys.PtrSize:
		lenmem = uintptr(old.len) * sys.PtrSize
		newlenmem = uintptr(cap) * sys.PtrSize
		capmem = roundupsize(uintptr(newcap) * sys.PtrSize)
		overflow = uintptr(newcap) > maxAlloc/sys.PtrSize
		newcap = int(capmem / sys.PtrSize)
	case isPowerOfTwo(et.size):
		var shift uintptr
		if sys.PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
		}
		lenmem = uintptr(old.len) << shift
		newlenmem = uintptr(cap) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
	default:
		lenmem = uintptr(old.len) * et.size
		newlenmem = uintptr(cap) * et.size
		capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.size)
	}

	// The check of overflow in addition to capmem > maxAlloc is needed
	// to prevent an overflow which can be used to trigger a segfault
	// on 32bit architectures with this example program:
	//
	// type T [1<<27 + 1]int64
	//
	// var d T
	// var s []T
	//
	// func main() {
	//   s = append(s, d, d, d, d)
	//   print(len(s), "\n")
	// }
	if overflow || capmem > maxAlloc {
		panic(errorString("growslice: cap out of range"))
	}

	// 创建一个新的底层数组指针
	var p unsafe.Pointer
	// 给p分配内存,增加容量
	if et.kind&kindNoPointers != 0 {
		p = mallocgc(capmem, nil, false)
		// The append() that calls growslice is going to overwrite from old.len to cap (which will be the new length).
		// Only clear the part that will not be overwritten.
		memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
	} else {
		// Note: can't use rawmem (which avoids zeroing of memory), because then GC can scan uninitialized memory.
		p = mallocgc(capmem, et, true)
		if writeBarrier.enabled {
			// Only shade the pointers in old.array since we know the destination slice p
			// only contains nil pointers because it has been cleared during alloc.
			bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(old.array), lenmem)
		}
	}
	// 将旧的数据copy给新的底层数组
	memmove(p, old.array, lenmem)

	return slice{p, old.len, newcap}
}

```
注意switch-cash那段代码中的`roundupsize`函数
```go
capmem = roundupsize(uintptr(newcap) * sys.PtrSize)

// Returns size of the memory block that mallocgc will allocate if you ask for the size.
func roundupsize(size uintptr) uintptr {
	if size < _MaxSmallSize {
		if size <= smallSizeMax-8 {
			// TODO 内存对齐在这里
			return uintptr(class_to_size[size_to_class8[(size+smallSizeDiv-1)/smallSizeDiv]])
		} else {
			return uintptr(class_to_size[size_to_class128[(size-smallSizeMax+largeSizeDiv-1)/largeSizeDiv]])
		}
	}
	if size+_PageSize < size {
		return size
	}
	return round(size, _PageSize)
}

// 将n四舍五入为a的倍数,a必须为2的幂
// round n up to a multiple of a.  a must be a power of 2.
func round(n, a uintptr) uintptr {
	return (n + a - 1) &^ (a - 1)
}
```
上述结果将导致最终得到的cap大于旧的cap的两倍或者一又四分之一.有些资料称这是为了**内存对齐**.

#### 大概流程:
- 如果指定的cap大于旧的slice的cap的两倍,那么新的cap的大小为指定的cap  
- 如果指定的cap小于旧的slice的cap的两倍  
    - 如果旧的slice的长度小于1024,那么新的cap为旧的cap的两倍  
    - 如果旧的slice的长度大于1024,则新的cap每次增加旧的cap的四分之一,直到接近指定的cap
- 获得新的cap后,计算该slice是否会内存溢出,并计算出旧的len占用的内存大小,新的len占用的内存大小,新的cap占用的内存大小
- 计算出来的新的cap因为"内存对齐",也许会大于"预期"
- 创建一个新的底层数组指针p,并根据新的cap内存大小来分配内存
- 将旧的slice中底层数组中的数据copy给新的底层数组
- 组成新的slice:新的底层数组,旧的slice的长度,新的cap大小

#### 遗留问题:
1. 函数的参数cap是什么时候传入的
2. 为什么为nil的slice可以直接进行append

#### 参考文章:
1. [码农桃花源-深度解密Go语言之Slice](https://www.cnblogs.com/qcrao-2018/p/10631989.html#%E4%B8%BA%E4%BB%80%E4%B9%88-nil-slice-%E5%8F%AF%E4%BB%A5%E7%9B%B4%E6%8E%A5-append)