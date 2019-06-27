# golang中的TDD
> TDD,即test driven development，测试驱动开发
[原文](https://medium.com/@pierreprinetti/test-driven-development-in-go-baeab5adb468)

**三条规则**
1. 直到完成测试代码，否则你不能写任何产品代码。
2. 直到一段代码不能满足测试需求，否则你不能更改任何生产代码。
3. 直到一段代码通过单元测试，否则你就不能再编写这段代码。

## 主要因素
### 第一步：定义API
我们的任务是找到一个整数的最小因子，即构成一个整数的所有素数。定义函数如下：
`func factorsOf(n int) []int`

然后我们定义检查函数`checkFunc`。它作为一个匹配函数能够获取`factorsOf`返回的所有数据

因为`factorsOf`只返回一个整数类型的切片，所以我们的检查类型只有整数类型切片。即
`type checkFunc func([]int) error`

### 第二步：第一个测试
这里我们使用**表驱动测试**
```go
func TestFactorsOf(t *testing.T) {
	type checkFunc func([]int) error

	// SECTION 1: checkers
	isEmptyList := func(have []int) error {
		if len(have) > 0 {
			return fmt.Errorf("Expected empty list, found %v.", have)
		}
		return nil
	}

	// SECTION 2: test cases
	tests := [...]struct {
		in    int
		check checkFunc
	}{
		// the first test: factors of 1
		{1, isEmptyList},
	}

	// SECTION 3: test logic
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Factors of %d", tc.in), func(t *testing.T) {
			factors := factorsOf(tc.in)
			if err := tc.check(factors); err != nil {
				t.Error(err)
			}
		})
	}
}
```
`isEmptyList `是我们第一个检查函数，如果我们的得到的数据不为空，它会返回错误。

`tests `是测试用例的数组，每一个用例都是一个匿名的结构体，结构体中有我们想要传递的参数和返回的检查函数，来检查是否负责规则。  
第一个测试用例表示，当我们输入1时，我们期待返回空切片。
  
在最后，我们遍历`tests`，并调用`t.Run`来检查每个用例是否得到了理想结果。

一个满足上边要求的生产代码可以是这样的。
```go

func factorsOf(n int) []int {
	return []int{}
}
```
> 在这里我们做到了，在测试代码完成前，不写成产代码。

### 第三步：迭代
第二个用例我们想测试一个非空的集合。
```go
is := func(want ...int) checkFunc {
  return func(have []int) error {
   if !reflect.DeepEqual(have, want) {
    return fmt.Errorf("Expected list %v, found %v.", want, have)
   }
   return nil
  }
 }
```
它是一个闭包函数，实际的匹配函数是基于期望的值动态生成的通过父类函数传递过来的。
```go
func TestFactorsOf(t *testing.T) {
	type checkFunc func([]int) error

	// SECTION 1: checkers
	isEmptyList := func(have []int) error {
		if len(have) > 0 {
			return fmt.Errorf("Expected empty list, found %v.", have)
		}
		return nil
	}
	is := func(want ...int) checkFunc {
		return func(have []int) error {
			if !reflect.DeepEqual(have, want) {
				return fmt.Errorf("Expected list %v, found %v.", want, have)
			}
			return nil
		}
	}

	// SECTION 2: test cases
	tests := [...]struct {
		in    int
		check checkFunc
	}{
		{1, isEmptyList},
		{2, is(2)},
	}

	// SECTION 3: test logic
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Factors of %d", tc.in), func(t *testing.T) {
			factors := factorsOf(tc.in)
			if err := tc.check(factors); err != nil {
				t.Error(err)
			}
		})
	}
}
```
```go
func factorsOf(n int) []int {
	var factors []int
	if n > 1 {
		factors = append(factors, 2)
	}
	return factors
}
```
可能上边的代码过于简单，但并不是作弊。我们应该提供符合最低要求的代码。  
接下来，我们继续增加用例
`
{3, is(3)}
`
产品代码让不需要改动，继续增加用例
`
{4, is(2, 2)}
`
然后改动我们的代码
```go

func factorsOf(n int) []int {
	var factors []int
	
	// return an empty slice if n <= 1
	if n > 1 {
		
		// factor 2
		if n%2 == 0 {
			factors = append(factors, 2)
			n /= 2
		}
		
		// append the remainder
		if n > 1 {
			factors = append(factors, n)
		}
	}
	
	return factors
}
```
继续增加用例
`
{5, is(5)},
{6, is(2, 3)},
{7, is(7)},
`
上边代码仍可以通过，继续增加用例
`{8, is(2, 2, 2)},`
代码不能通过了，再次改变我们的代码
```go
func factorsOf(n int) []int {
	var factors []int
	if n > 1 {
		
		// factor all the 2s
		for n%2 == 0 {
			factors = append(factors, 2)
			n /= 2
		}
		
		// append the remainder
		if n > 1 {
			factors = append(factors, n)
		}
	}
	return factors
}
```
继续增加用例
`{9, is(3, 3)},`
然后发现已有的代码不能满足上述用例，所以需要更改代码
```go
func factorsOf(n int) []int {
	var factors []int
	for divisor := 2; n > 1; divisor++ {
		for n%divisor == 0 {
			factors = append(factors, divisor)
			n /= divisor
		}
	}
	return factors
}
```

> 由上边的过程，我们可以得出结论：测试驱动开发就是一个使用不断增加解决办法的方式来解决问题。
就像 Mr. Martin说的：**我们没有设计任何算法，那它从哪里来的呢？我们只是一个个通过测试用例**

*最终的代码在`tdd_test.go`中可以找到。*

感受：**虽然结果很完美，但是实际开发中，如果每个用例都这么写，所花费的时间应该很长，所以一次性考虑好所有的可能情景，然后写好测试。然后一次性写出生产代码并测试。这样应该会好一点。**
