源码位置：`golang.org/x/time/rate`

`go version go1.13.5 windows/amd64`

实现作用：**限制访问频率**

实现机制类似与`令牌桶`

主要数据结构：

```go
type Limiter struct {
	limit Limit // 1秒内允许访问的次数
	burst int // 一次访问的做多的事件个数，

	mu     sync.Mutex
	tokens float64 // 当前能使用的token的个数

	last time.Time // 上一次tokens字段更新的事件
	lastEvent time.Time // 事件发生最新的时间 (可能是过去，也可以是未来，因为可以预约)
}
```

`limit`是`float64`类型，标识。也可以通过`Every`来生成

```go
type Limit float64
// 通过最小的时间间隔来生成limit
func Every(interval time.Duration) Limit {
	if interval <= 0 {
		return Inf
	}
	return 1 / Limit(interval.Seconds())
}
```

什么是**事件**呢？在这里一次访问可能有多个事件，每个事件都要消耗一个`token`

第二重要的数据结构

```go
// 有些事件会预约token，这个事件的信息就保存在Reservation 。当然预约也能被limitter取消
type Reservation struct {
	ok        bool
	lim       *Limiter
	tokens    int
	timeToAct time.Time
	limit Limit
}
```



先看一个很基础的方法。

```go
// 通过传入时间、事件个数和最大的等待时间来获得预约信息
func (lim *Limiter) reserveN(now time.Time, n int, maxFutureReserve time.Duration) Reservation {
	lim.mu.Lock()
	// 如果没有限制，直接返回一个立即就能获取到想要的token个数的预约信息
	if lim.limit == Inf {
		lim.mu.Unlock()
		return Reservation{
			ok:        true,
			lim:       lim,
			tokens:    n,
			timeToAct: now,
		}
	}

    // 获取上次token访问时间以及当前能够使用token的数量，now还是那个now（所以为啥要返回呢）
	now, last, tokens := lim.advance(now)

	// 计算访问后剩余的token数量
	tokens -= float64(n)

	// 计算需要等待所需token的时间
	var waitDuration time.Duration
	if tokens < 0 {
		waitDuration = lim.limit.durationFromTokens(-tokens)
	}

	// 计算能否在最长等待时间前得到所需token
	ok := n <= lim.burst && waitDuration <= maxFutureReserve

	// 准备预约
	r := Reservation{
		ok:    ok,
		lim:   lim,
		limit: lim.limit,
	}
	if ok {
		r.tokens = n
		r.timeToAct = now.Add(waitDuration)
	}

	// 更新状态
	if ok {
		lim.last = now
		lim.tokens = tokens
		lim.lastEvent = r.timeToAct
	} else {
		lim.last = last
	}

	lim.mu.Unlock()
	return r
}
```

上个方法中用到的另一个方法： 根据现在的时间，获取到`limitter`最新访问的时间、上次访问的时间与最多的能够使用的`token`数量

```go
// 根据现在的时间，获取到limitter最新访问的时间、上次访问的时间与最多的能够使用的token数量
func (lim *Limiter) advance(now time.Time) (newNow time.Time, newLast time.Time, newTokens float64) {
    // 获取上次token获取的时间，如果上次时间晚于现在，说明有预支，那么将当前时间设置为上次访问时间，那么下面根据时间间隔计算能够使用的token的数量时，会得到0，因为已知被预支了。
	last := lim.last
	if now.Before(last) {
		last = now
	}

	// 通过一次允许做多消耗的token个数与当前剩余的token 个数，来计算能够使用的最多token对应的时间间隔.防止因为上次访问过长导致分配过多的token
	maxElapsed := lim.limit.durationFromTokens(float64(lim.burst) - lim.tokens)
	elapsed := now.Sub(last)
	if elapsed > maxElapsed {
		elapsed = maxElapsed
	}

	// 根据时间间隔来获取能够使用的token个数
	delta := lim.limit.tokensFromDuration(elapsed)
	tokens := lim.tokens + delta
	if burst := float64(lim.burst); tokens > burst {
		tokens = burst
	}

	return now, last, tokens
}
```

`reserveN`被三个方法调用，分别时：

```go
// 当前能否获取到n个token
func (lim *Limiter) AllowN(now time.Time, n int) bool {
	return lim.reserveN(now, n, 0).ok
}
// 当前获取n个token需要的等待信息
func (lim *Limiter) ReserveN(now time.Time, n int) *Reservation {
	r := lim.reserveN(now, n, InfDuration)
	return &r
}
// 阻塞直到获取到n个token
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error) {
    // ....
}
```

代码中最核心的三个方法也是这三个。