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
    
实现原理