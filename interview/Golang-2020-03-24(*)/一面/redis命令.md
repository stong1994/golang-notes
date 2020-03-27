1. sorted set：
    1. 取前十条热门微博
        > zrange hot 0 9
    2. 取第九条热门数据
        > zrange hot 8 8 
2. hash
    1. 获取所有的field
        > hkeys key
    2. 获取所有的value
        > hvals key
3. set
    1. 获取多少个值
        > scard key
    2. 删除一个值
        > srem key field
    
[redis功能文档](http://redisdoc.com/topic/index.html)

##### 命令总结

> 1. 五个基本类型：string、list、hash、set、sorted set
> 2. 每个命令的首字母基本都是该类型的首字母（string和sorted set除外，sorted set的s已经被set抢占，因此用z表示）
> 3. 如果操作多个key，命令以m开头（multi-），如 mset、mget
> 4. 如果给key的value设置属性，则把这个属性放在key value中间，如SETEX key expire-time value 

#### 基础命令

| 命令                                      | 解释                                                         |
| ----------------------------------------- | ------------------------------------------------------------ |
| EXISTS exists-key                         | 是否存在key                                                  |
| DEL key [key …]                           | 删除key                                                      |
| EXPIRE key seconds                        | 为给定 `key` 设置生存时间，以秒为单位                        |
| PEXPIRE key milliseconds                  | 与上相似，生存时间以毫秒为单位                               |
| EXPIREAT key timestamp                    | 为给定 `key` 设置生存时间，时间为UNIX 时间戳                 |
| PEXPIREAT key milliseconds-timestamp      | 这个命令和 `expireat` 命令类似，但它以毫秒为单位设置 `key` 的过期 unix 时间戳 |
| TTL key                                   | 查看秒级剩余过期时间                                         |
| PTTL key                                  | 查看毫秒级剩余过期时间                                       |
| PERSIST key                               | 移除给定 `key` 的生存时间                                    |
|                                           |                                                              |
| TYPE key                                  | 返回 `key` 所储存的值的类型。                                |
| RENAME key newkey                         | 将 `key` 改名为 `newkey` 。                                  |
| RENAMENX key newkey                       | 当且仅当 `newkey` 不存在时，将 `key` 改名为 `newkey` 。      |
| MOVE key db                               | 将当前数据库的 `key` 移动到给定的数据库 `db` 当中。          |
| RANDOMKEY                                 | 从当前数据库中随机返回(不删除)一个 `key` 。                  |
| DBSIZE                                    | 返回当前数据库的 key 的数量。                                |
| KEYS pattern                              | 查找所有符合给定模式 `pattern` 的 `key`                      |
| SCAN cursor [MATCH pattern] [COUNT count] | `SCAN` 命令及其相关的 `SSCAN` 命令、 `HSCAN` 命令和 `ZSCAN` 命令都用于增量地迭代 |
| FLUSHDB                                   | 清空当前数据库中的所有 key。                                 |
| FLUSHALL                                  | 清空整个 Redis 服务器的数据(删除所有数据库的所有 key )。     |
| SELECT index                              | 切换到指定的数据库，数据库索引号 `index` 用数字值指定，以 `0` 作为起始索引值。 |
| SWAPDB db1 db2                            | 对换指定的两个数据库， 使得两个数据库的数据立即互换。        |
| SORT                                      |                                                              |
| MONITOR                                   | 实时显示redis服务器执行的命令                                |

#### 字符串

| 命令                           | 描述                                                         |
| ------------------------------ | ------------------------------------------------------------ |
| SET key "value"                |                                                              |
| SET key value EX expire-time   | 秒级过期时间                                                 |
| TTL key-with-expire-time       | 查看秒级剩余过期时间                                         |
| SET key"moto" PX expire-time   | 毫秒级过期时间                                               |
| PTTL key-with-pexpire-time     | 查看毫秒级剩余过期时间                                       |
| SET not-exists-key value NX    | 键不存在时才能设置值                                         |
| SET exists-key "new-value" XX  | 键存在时才能设置值                                           |
| SETEX key expire-time value    | 在一个原子操作中同时设置值和过期时间                         |
| PSETEX key expire-time value   | 接上，过期时间为毫秒                                         |
| GET key                        | 获取字符串类型的值                                           |
| GETSET key value               | 将键 `key` 的值设为 `value` ， 并返回键 `key` 在被设置之前的旧值 |
| STRLEN key                     | 返回键 `key` 储存的字符串值的长度。                          |
| APPEND key value               | 如果键 `key` 已经存在并且它的值是一个字符串， `APPEND` 命令将把 `value` 追加到键 `key` 现有值的末尾。如果 `key` 不存在， `APPEND` 就简单地将键 `key` 的值设为 `value` ， 就像执行 `SET key value` 一样。返回字符串长度。 |
| SETRANGE key offset value      | 从偏移量 `offset` 开始， 用 `value` 参数覆写(overwrite)键 `key` 储存的字符串值。 |
| GETRANGE key start end         | 返回键 `key` 储存的字符串值的指定部分， 字符串的截取范围由 `start` 和 `end` 两个偏移量决定 (包括 `start` 和 `end` 在内)。 |
| INCR key                       | 为键 `key` 储存的数字值加上一。                              |
| INCRBY key increment           | 为键 `key` 储存的数字值加上增量 `increment` 。               |
| INCRBYFLOAT key increment      | 为键 `key` 储存的值加上浮点数增量 `increment` 。             |
| DECR key                       | 为键 `key` 储存的数字值减去一。                              |
| DECRBY key decrement           | 将键 `key` 储存的整数值减去减量 `decrement` 。               |
| MSET key value [key value …]   | 同时为多个键设置值。如果某个给定键已经存在， 那么 `MSET` 将使用新值去覆盖旧值 |
| MSETNX key value [key value …] | 当且仅当所有给定键都不存在时， 为所有给定键设置值。          |
| MGET key [key …]               | 返回给定的一个或多个字符串键的值。如果给定的字符串键里面， 有某个键不存在， 那么这个键的值将以特殊值 `nil` 表示。 |

#### HASH:

| 命令                                           | 解释                                                         |
| ---------------------------------------------- | ------------------------------------------------------------ |
| HSET hash field value                          | 将哈希表 `hash` 中域 `field` 的值设置为 `value` 。新建field返回1，覆盖旧值返回0 |
| HSETNX hash field value                        | 当且仅当域 `field` 尚未存在于哈希表的情况下， 将它的值设置为 `value` 。如果哈希表 `hash` 不存在， 那么一个新的哈希表将被创建并执行 `HSETNX` 命令。 |
| HGET hash field                                | `HGET` 命令在默认情况下返回给定域的值。                      |
| HEXISTS hash field                             | `HEXISTS` 命令在给定域存在时返回 `1` ， 在给定域不存在时返回 `0` 。 |
| HDEL key field [field …]                       | 删除哈希表 `key` 中的一个或多个指定域，不存在的域将被忽略。返回删除的`field`的个数 |
| HLEN                                           | 返回哈希表 `key` 中域的数量                                  |
| HSTRLEN key field                              | 返回哈希表 `key` 中， 与给定域 `field` 相关联的值的字符串长度（string length）。 |
| HINCRBY key field increment                    | 为哈希表 `key` 中的域 `field` 的值加上增量 `increment` 。    |
| HINCRBYFLOAT key field increment               | 为哈希表 `key` 中的域 `field` 加上浮点数增量 `increment` 。  |
| HMSET key field value [field value …]          | 同时将多个 `field-value` (域-值)对设置到哈希表 `key` 中。    |
| HMGET key field [field …]                      | 返回哈希表 `key` 中，一个或多个给定域的值。                  |
| HKEYS key                                      | 返回哈希表 `key` 中的所有域。                                |
| HVALS key                                      | 返回哈希表 `key` 中所有域的值。                              |
| HGETALL key                                    | 返回哈希表 `key` 中，所有的域和值。                          |
| HSCAN key cursor [MATCH pattern] [COUNT count] |                                                              |

#### LIST

| 命令                                  | 描述                                                         |
| ------------------------------------- | ------------------------------------------------------------ |
| LPUSH key value [value …]             | 将一个或多个值 `value` 插入到列表 `key` 的表头               |
| LPUSHX key value                      | 将值 `value` 插入到列表 `key` 的表头，当且仅当 `key` 存在并且是一个列表。返回`list`的长度 |
| RPUSH key value [value …]             | 将一个或多个值 `value` 插入到列表 `key` 的表尾(最右边)。     |
| RPUSHX key value                      | 将值 `value` 插入到列表 `key` 的表尾，当且仅当 `key` 存在并且是一个列表。 |
| LPOP key                              | 移除并返回列表 `key` 的头元素。                              |
| RPOP key                              | 移除并返回列表 `key` 的尾元素。                              |
| RPOPLPUSH source destination          | `RPOP source` + `LPUSH destination`                          |
| LREM key count value                  | 根据参数 `count` 的值，移除列表中与参数 `value` 相等的元素。`count> 0`:从表头开始向表尾搜索，移除与 `value` 相等的元素，数量为 `count`;`count < 0` : 从表尾开始向表头搜索，移除与 `value` 相等的元素，数量为 `count` 的绝对值;`count = 0` : 移除表中所有与 `value` 相等的值 |
| LLEN key                              | 返回列表 `key` 的长度。                                      |
| LINDEX key index                      | 返回列表 `key` 中，下标为 `index` 的元素。                   |
| LINSERT key BEFORE\|AFTER pivot value | 将值 `value` 插入到列表 `key` 当中，位于值 `pivot` 之前或之后。`key`或者`pivot`不存在时，不执行任何操作 |
| LSET key index value                  | 将列表 `key` 下标为 `index` 的元素的值设置为 `value` 。      |
| LRANGE key start stop                 | 返回列表 `key` 中指定区间内的元素，区间以偏移量 `start` 和 `stop` 指定。 |
| LTRIM key start stop                  | 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。 |
| BLPOP key [key …] timeout             | 它是 `LPOP key` 命令的阻塞版本，当给定列表内没有任何元素可供弹出的时候，连接将被 `BLPOP` 命令阻塞，直到等待超时或发现可弹出元素为止;当给定多个 `key` 参数时，按参数 `key` 的先后顺序依次检查各个列表，弹出第一个非空列表的头元素。`timeout`为0代表永久阻塞直到弹出元素 |
| BRPOP key [key …] timeout             | 与上方向相反                                                 |
| BRPOPLPUSH source destination timeout | `BRPOPLPUSH` 是 `RPOPLPUSH source destination` 的阻塞版本，当给定列表 `source` 不为空时， `BRPOPLPUSH`的表现和 `RPOPLPUSH source destination` 一样。 |
|                                       |                                                              |
|                                       |                                                              |

#### SET

| 命令                                           | 描述                                                         |
| ---------------------------------------------- | ------------------------------------------------------------ |
| SADD key member [member …]                     | 将一个或多个 `member` 元素加入到集合 `key` 当中，已经存在于集合的 `member` 元素将被忽略。 |
| SISMEMBER key member                           | 判断 `member` 元素是否集合 `key` 的成员。                    |
| SPOP key                                       | 移除并返回集合中的一个随机元素。                             |
| SRANDMEMBER key [count]                        | 如果 `count` 为正数，且小于集合基数，那么命令返回一个包含 `count` 个元素的数组，数组中的元素**各不相同**。如果 `count` 大于等于集合基数，那么返回整个集合。 如果 `count` 为负数，那么命令返回一个数组，数组中的元素**可能会重复出现多次**，而数组的长度为 `count` 的绝对值。 |
| SREM key member [member …]                     | 移除集合 `key` 中的一个或多个 `member` 元素，不存在的 `member` 元素会被忽略。 |
| SMOVE source destination member                | 将 `member` 元素从 `source` 集合移动到 `destination` 集合。  |
| SCARD key                                      | 返回集合 `key` 的基数(集合中元素的数量)。                    |
| SMEMBERS key                                   | 返回集合 `key` 中的所有成员。                                |
| SSCAN key cursor [MATCH pattern] [COUNT count] |                                                              |
| SINTER key [key …]                             | 返回一个集合的全部成员，该集合是所有给定集合的交集。         |
| SINTERSTORE destination key [key …]            | 接上，将结果保存到 `destination` 集合，而不是简单地返回结果集。 |
| SUNION key [key …]                             | 返回一个集合的全部成员，该集合是所有给定集合的并集。         |
| SUNIONSTORE destination key [key …]            | 接上，将结果保存到 `destination` 集合，而不是简单地返回结果集。 |
| SDIFF key [key …]                              | 返回一个集合的全部成员，该集合是所有给定集合之间的差集。     |
| SDIFFSTORE destination key [key …]             | 街上，将结果保存到 `destination` 集合，而不是简单地返回结果集。 |

#### SORTED SET

| 命令                                                         | 描述                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| ZADD key score member [[score member] [score member] …]      | 将一个或多个 `member` 元素及其 `score` 值加入到有序集 `key` 当中。 |
| ZSCORE key member                                            | 返回有序集 `key` 中，成员 `member` 的 `score` 值。           |
| ZINCRBY key increment member                                 | 为有序集 `key` 的成员 `member` 的 `score` 值加上增量 `increment` 。 |
| ZCARD key                                                    | 返回有序集 `key` 的基数。                                    |
| ZCOUNT key min max                                           | 返回有序集 `key` 中， `score` 值在 `min` 和 `max` 之间(默认包括 `score` 值等于 `min` 或 `max` )的成员的数量。 |
| ZRANGE key start stop [WITHSCORES]                           | 有序集 `key` 中，成员的位置按 `score` 值递增(从小到大)来排序返回。指定`WITHSCORES`将使成员与分数一起返回。 |
| ZREVRANGE key start stop [WITHSCORES]                        | 与上排序方向相反                                             |
| ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]  | 返回有序集 `key` 中，所有 `score` 值介于 `min` 和 `max` 之间(包括等于 `min` 或 `max` )的成员。有序集成员按 `score` 值递增(从小到大)次序排列。`min` 和 `max` 可以是 `-inf` 和 `+inf`,在不知道最大/小值时使用。可以通过给参数前增加 `(` 符号来使用可选的`开区间` (小于或大于)。 |
| ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count] | 与上排序方向相反                                             |
| ZRANK key member                                             | 返回有序集 `key` 中成员 `member` 的排名。其中有序集成员按 `score` 值递增(从小到大)顺序排列。 |
| ZREVRANK key member                                          | 与上排序方向相反                                             |
| ZREM key member [member …]                                   | 移除有序集 `key` 中的一个或多个成员，不存在的成员将被忽略。  |
| ZREMRANGEBYRANK key start stop                               | 移除有序集 `key` 中，指定排名(rank)区间内的所有成员。        |
| ZREMRANGEBYSCORE key min max                                 | 移除有序集 `key` 中，所有 `score` 值介于 `min` 和 `max` 之间(包括等于 `min` 或 `max` )的成员。 |
| ZRANGEBYLEX key min max [LIMIT offset count]                 | 当有序集合的所有成员都具有相同的分值时， 有序集合的元素会根据成员的字典序来进行排序， 而这个命令则可以返回给定的有序集合键 `key` 中， 值介于 `min` 和 `max` 之间的成员。`min`和`max`的表示可以用`-`与`+`表示最小与最大，区间必须用`(`或`[`表示开还是闭 |
| ZLEXCOUNT key min max                                        | 对于一个所有成员的分值都相同的有序集合键 `key` 来说， 这个命令会返回该集合中， 成员介于 `min` 和 `max` 范围内的元素数量。 |
| ZREMRANGEBYLEX key min max                                   | 对于一个所有成员的分值都相同的有序集合键 `key` 来说， 这个命令会移除该集合中， 成员介于 `min` 和 `max` 范围内的所有元素。 |
| ZSCAN key cursor [MATCH pattern] [COUNT count]               |                                                              |
| ZUNIONSTORE destination numkeys key [key …] [WEIGHTS weight [weight …]] [AGGREGATE SUM\|MIN\|MAX] | 计算给定的一个或多个有序集的并集，其中给定 `key` 的数量必须以 `numkeys` 参数指定，并将该并集(结果集)储存到 `destination` 。使用 `WEIGHTS` 选项，你可以为 *每个* 给定有序集 *分别* 指定一个乘法因子，每个给定有序集的所有成员的 `score` 值在传递给聚合函数之前都要先乘以该有序集的因子。 |
| ZINTERSTORE destination numkeys key [key …] [WEIGHTS weight [weight …]] [AGGREGATE SUM\|MIN\|MAX] | 与上相似，由并集改为了交集                                   |

#### 事务

| 命令              | 解释                                                         |
| ----------------- | ------------------------------------------------------------ |
| MULTI             | 标记事务开始                                                 |
| EXEC              | 执行所有事务块内的命令。                                     |
| DISCARD           | 取消所有事务，在取消事务的同时也会取消所有对 key 的监视      |
| WATCH key [key …] | 监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断。 |
| UNWATCH           | 取消 WATCH 命令对所有 key 的监视。                           |

#### 持久化

| 命令         | 解释                                                         |
| ------------ | ------------------------------------------------------------ |
| SAVE         | `SAVE` 命令执行一个同步保存操作，将当前 Redis 实例的所有数据快照以` RDB` 文件的形式保存到硬盘。它会阻塞所有客户端，所以很少使用。 |
| BGSAVE       | 在后台异步保存当前数据库的数据到磁盘。`BGSAVE` 命令执行之后立即返回 `OK` ，然后 Redis fork 出一个新子进程，原来的 Redis 进程(父进程)继续处理客户端请求，而子进程则负责将数据保存到磁盘，然后退出。 |
| BGREWRITEAOF | 执行一个 `AOF文件`重写操作。重写会创建一个当前 AOF 文件的体积优化版本。 |
| LASTSAVE     | 返回最近一次 Redis 成功将数据保存到磁盘上的时间，以 UNIX 时间戳格式表示。 |

