# go测试的一种模式

[原文链接](https://medium.com/@pierreprinetti/a-pattern-for-go-tests-3468b51535)

我一直在思考如何写好go的测试，很多博客都提到了引用外部的第三方库，比如`github.com/golang/mock`，通过这些工具能够加速开发速度。
但是我认为我们应该拥抱GO的真正的具有优势的特性，而不是直接利用别人封装好的。有良好的编码习惯，不仅能够提供代码质量，还能增强洞察力。

于是，我开始查看GO的核心库，并且找到了这么一个用于测试目的的包 [net/http/httptest](https://golang.org/pkg/net/http/httptest/)

这里有个例子`recorder_test.go`：
```go
func TestRecorder(t *testing.T) {
	type checkFunc func(*ResponseRecorder) error
	check := func(fns ...checkFunc) []checkFunc { return fns }

	hasStatus := func(want int) checkFunc {
		return func(rec *ResponseRecorder) error {
			if rec.Code != want {
				return fmt.Errorf("expected status %d, found %d", want, rec.Code)
			}
			return nil
		}
	}
	hasContents := func(want string) checkFunc {
		return func(rec *ResponseRecorder) error {
			if have := rec.Body.String(); have != want {
				return fmt.Errorf("expected body %q, found %q", want, have)
			}
			return nil
		}
	}
	hasHeader := func(key, want string) checkFunc {
		return func(rec *ResponseRecorder) error {
			if have := rec.Result().Header.Get(key); have != want {
				return fmt.Errorf("expected header %s: %q, found %q", key, want, have)
			}
			return nil
		}
	}

	tests := [...]struct {
		name   string
		h      func(w http.ResponseWriter, r *http.Request)
		checks []checkFunc
	}{
		{
			"200 default",
			func(w http.ResponseWriter, r *http.Request) {},
			check(hasStatus(200), hasContents("")),
		},
		{
			"first code only",
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(201)
				w.WriteHeader(202)
				w.Write([]byte("hi"))
			},
			check(hasStatus(201), hasContents("hi")),
		},
		{
			"write string",
			func(w http.ResponseWriter, r *http.Request) {
				io.WriteString(w, "hi first")
			},
			check(
				hasStatus(200),
				hasContents("hi first"),
				hasHeader("Content-Type", "text/plain; charset=utf-8"),
			),
		},
	}

	r, _ := http.NewRequest("GET", "http://foo.com/", nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := http.HandlerFunc(tt.h)
			rec := NewRecorder()
			h.ServeHTTP(rec, r)
			for _, check := range tt.checks {
				if err := check(rec); err != nil {
					t.Error(err)
				}
			}
		})
	}
}
```

代码分为三部分：
#### 第一部分
1. 定义检查函数`checkFunc`。这个函数的参数为我们想要测试的每个值。在例子中，我们只用获取`ResponseRecorder`就能得到所有返回值，所以`ResponseRecorder`就是检查函数的参数。
2. 创建匹配函数。例子中为`hasStatus`，`hasContents `,`hasHeader `。这些函数是闭包的，它们提供我们期待的结果，并且通过返回检查函数来检查是否能得到我们想要的结果。如果不是，则报错。

#### 第二部分
1. 用匿名的结构体来保存测试数据，例子中为`tests `，结构体定义为
    1. 这个用例的描述
    2. 输入
    3. 一个检查函数组成的切片——通过调用匹配函数得到的检查函数。

#### 第三部分
这里我们来检查我们的错误。使用测试用例来执行我们的代码，一旦有错误产生，就调用`t.Error()`

