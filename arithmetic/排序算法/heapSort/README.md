# 堆排序
> 文中代码实现为小顶堆
>
> 用途：处理一般的静态队列（FIFO/LIFO）所不能完成的任务，用于动态队列，即优先队列。如医院看病，并非只是FIFO，退役军人等具有优先级的资格。  
  代码取自：https://github.com/pingcap-incubator/tinykv  
> Pop(), 获取最小元素，并删除，由于是最小堆，那么获取的是首元素
  Remove(),获取最后一个元素，由于是最大堆，那么获取的是尾元素  
> 堆排序介绍：https://www.cnblogs.com/chengxiao/p/6129630.html