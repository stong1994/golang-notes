1. tcp的三次握手
    1. 第一次握手(SYN=1, seq=x):  
      客户端发送一个 TCP 的 SYN 标志位置1的包，指明客户端打算连接的服务器的端口，以及初始序号 X,保存在包头的序列号(Sequence Number)字段里。
      发送完毕后，客户端进入 SYN_SEND 状态。
    2. 第二次握手(SYN=1, ACK=1, seq=y, ACKnum=x+1):  
      服务器发回确认包(ACK)应答。即 SYN 标志位和 ACK 标志位均为1。服务器端选择自己 ISN 序列号，放到 Seq 域里，同时将确认序号(Acknowledgement Number)设置为客户的 ISN 加1，即X+1。 发送完毕后，服务器端进入 SYN_RCVD 状态。
    3. 第三次握手(ACK=1，ACKnum=y+1)  
      客户端再次发送确认包(ACK)，SYN 标志位为0，ACK 标志位为1，并且把服务器发来 ACK 的序号字段+1，放在确定字段中发送给对方，并且在数据段放写ISN的+1  
      发送完毕后，客户端进入 ESTABLISHED 状态，当服务器端接收到这个包时，也进入 ESTABLISHED 状态，TCP 握手结束。
2. tcp的四次挥手
    1. 第一次挥手(FIN=1，seq=x)  
       假设客户端想要关闭连接，客户端发送一个 FIN 标志位置为1的包，表示自己已经没有数据可以发送了，但是仍然可以接受数据。
       发送完毕后，客户端进入 FIN_WAIT_1 状态。
    2. 第二次挥手(ACK=1，ACKnum=x+1)  
       服务器端确认客户端的 FIN 包，发送一个确认包，表明自己接受到了客户端关闭连接的请求，但还没有准备好关闭连接。
       发送完毕后，服务器端进入 CLOSE_WAIT 状态，客户端接收到这个确认包之后，进入 FIN_WAIT_2 状态，等待服务器端关闭连接。
    3. 第三次挥手(FIN=1，seq=y)  
       服务器端准备好关闭连接时，向客户端发送结束连接请求，FIN 置为1。
       发送完毕后，服务器端进入 LAST_ACK 状态，等待来自客户端的最后一个ACK。
    4. 第四次挥手(ACK=1，ACKnum=y+1)  
       客户端接收到来自服务器端的关闭请求，发送一个确认包，并进入 TIME_WAIT状态，等待可能出现的要求重传的 ACK 包。
       服务器端接收到这个确认包之后，关闭连接，进入 CLOSED 状态。
       客户端等待了某个固定时间（两个最大段生命周期，2MSL，2 Maximum Segment Lifetime）之后，没有收到服务器端的 ACK ，认为服务器端已经正常关闭连接，于是自己也关闭连接，进入 CLOSED 状态。
3. time_wait发生在哪个阶段以及有什么作用  
   > time_wait是tcp释放连接的四次挥手后的主动关闭连接方的状态。  
   
   > 作用1：为了保证客户端发送的最后一个ack报文段能够到达服务器。因为这最后一个ack确认包可能会丢失，然后服务器就会超时重传第三次挥手的fin信息报，然后客户端再重传一次第四次挥手的ack报文。如果没有这2msl，客户端发送完最后一个ack数据报后直接关闭连接，那么就接收不到服务器超时重传的fin信息报，那么服务器就不能按正常步骤进入close状态。     
   
   > 作用2：在第四次挥手后，经过2msl的时间足以让本次连接产生的所有报文段都从网络中消失，这样下一次新的连接中就肯定不会出现旧连接的报文段了。
4. 为什么tcp是三次握手而不是两次握手？
    第三次握手时为了防止已失效的连接请求报文段有传送到B，因而产生错误。
    > 所谓“防止已失效的连接请求报文”是这样产生的。考虑一种正常情况。A发出连接请求，但因连接请求报文丢失而未收到确认。于是A再重传一次请求连接。后来收到了确认，建立了连接。数据传输完毕后，就释放了连接。A共发送了两个连接请求报文段。其中第一个丢失，第二个到达了B。没有“已失效的请求连接报文段”。  
    > 现假定出现一种异常情况，即A发出的第一个请求连接报文段并没有丢失，而是在某些网络结点长时间滞留了，以至到连接释放以后的某个时间才到达B。本来这是一个已失效的报文段。但B收到此失效的连接请求报文段后，就误认为是A又发出一次新的连接请求。于是就向A发出确认报文段，同意建立连接。假定不采用三次握手，那么只要B发出确认，新的连接就建立了。  
    > 由于现在A并没有发出建立请求的连接，因此不会理睬B的确认，也不会向B发送数据，但B却以为新的运输连接已经建立了，并一直等待A发来的数据。B的许多资源就这样白白浪费了。
   
#### 参考资料
1. [time_wait发生在哪个阶段以及有什么作用](https://zhuanlan.zhihu.com/p/51961509)
2. [为什么tcp是三次握手而不是两次握手](https://zhuanlan.zhihu.com/p/51448333)
3. [cp的三次握手和四次挥手 ACK SYN 第四次挥手后的等待2msl的原因](https://zhuanlan.zhihu.com/p/37641172)
4. [三次握手和四次挥手总结](https://juejin.im/post/5d9c284b518825095879e7a5)