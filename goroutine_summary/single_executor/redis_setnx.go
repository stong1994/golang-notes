package single_executor

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// redis本身就是单线程，并发安全的，再利用redis中命令setnx的特点（只能设置一次），来实现并发情况下的单一任务执行者。

// SETNX 是『SET if Not eXists』(如果不存在，则 SET)的简写。
// 		只在键 key 不存在的情况下， 将键 key 的值设置为 value 。
// 		若键 key 已经存在， 则 SETNX 命令不做任何动作。
/*
	redis> EXISTS job                # job 不存在
	(integer) 0

	redis> SETNX job "programmer"    # job 设置成功
	(integer) 1

	redis> SETNX job "code-farmer"   # 尝试覆盖 job ，失败
	(integer) 0

	redis> GET job                   # 没有被覆盖
	"programmer"
*/
var lockKey = "counter_lock"
var counterKey = "counter_key"

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123", // password set
		DB:       0,     // use default DB
	})
}

func incr() {
	client := redisClient()
	// lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	// unlock  todo 实现单一任务执行者需要注释掉下边代码
	/*	delResp := client.Del(lockKey)
		unlockSuccess, err := delResp.Result()
		if err == nil && unlockSuccess > 0 {
			println("unlock success!")
		} else {
			println("unlock failed", err)
		}*/
}
