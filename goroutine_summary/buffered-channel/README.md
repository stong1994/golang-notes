## 有缓冲的通道的用处

- 限制goroutine的数量

- 能够防止goroutine泄露，可以查看buffered-channel-complex/goroutine-lead_test.go

* 代码来自https://medium.com/capital-one-tech/buffered-channels-in-go-what-are-they-good-for-43703871828

* 代码有略微改动

- 构建“资源池”，避免重复的分配、释放。 [代码来源](https://www.kancloud.cn/kancloud/effective/72213)