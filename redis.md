## 什么是redis 
是基于内存的数据结构存储,用作数据库，缓存， 消息代理
## redis 有什么优缺点?
#### 优点
1. 完全基于内存，速度快
2. 有丰富的数据结构
3. 将数据持久化

#### 缺点
1. 单线程，无法利用多核
2. 内存数据库，有可能导致内存不够用

## redis 为什么这么快？
1. 完全基于内存，绝大部分请求都基于内存，非常快
2. 数据结构简单，对数据操作也简单
3. 采用单线程，避免了不必要的上下文切换和竞争条件
4. 使用IO 多路复用，非阻塞IO

## redis 有哪些数据类型？
string list hash set  sorted set


## 持久化
1. RDB (redis database): 基于时间点的快照
2. AOF(append only file) 写操作以日志记录下来


RDB 的优点：
1. RDB是个紧凑的时间点表示的单个文件，非常适合作备份。


## 缓存穿透
访问到数据库也没有的数据
## 缓存击穿
一个热点key失效，大量的访问打到数据库
## 缓存雪崩
是大面积版的缓存击穿，key在同一时间失效，大量的访问打到数据库。解决方案：给ttl 额外加上一个随机的时间


redis red key
算法：
1. 计算当前的时间戳
2. 依次访问每个redis,尝试set key




redis 高可用
哨兵模式
提供以下几点功能
- 监控 不断检查master 和slaver 是否正常运行
- 通知
- 故障转移
- 提供配置 client 可以请求哨兵拿到master的地址



## redis hash 使用的是什么数据结构？
    ziplist 和 hashtable

## redis 速度为什么这么快？
    1. 完全基于内存
    2. 单核cpu, 避免了上下文调度和竞争
    3. 采用非阻塞IO, IO 多路复用
    4. 
redis 使用场景？

## redis 的内存驱逐策略
 当内存达到限制时，会采用驱逐策略，策略有： (noeviction)不驱逐直接报错, LRU, random, TTL, redis-4.0 加了 LFU
 默认策略是 noeviction 查看命令 `config get maxmemory-policy`
 获取最大内存命令 `config get maxmemory` 将maxmemory 设置为0会导致没有内存限制，这是64位操作系统的默认设置。对于32位操作系统maxmemory默认是3GB.

 ## redis 集群（指的是分片 partation）
 ### redis 集群能够提供什么功能？
 1. 数据分片
 2. 在某些节点failed 后仍然能够操作

redis cluster 需要2个tcp 连接。默认是6379，在加10000，是16379。 16379用作节点之间的通信。

## redis cluster data sharding
    redis 的槽位是固定的，16384。 每个节点会被分配一些槽位。crc16(key)%16384

## redis哈希槽 和一致性哈希的区别？
1. redis hash槽并不是闭合的，它一共有16384个槽，使用CRC16算法计算key的hash值，与16384取模，确定数据在哪个槽中，从而找到所属的redis节点；一致性hash表示一个0到2^32的圆环，对数据计算hash后落到该圆环中，顺时针第一个节点为其所属服务。

2. 一致性hash是通过虚拟节点去避免服务节点宕机后数据转移造成的服务访问量激增、内存占用率过高、数据倾斜等问题，保证数据完整性和集群可用性的；而hash集群是使用主从节点的形式，主节点提供读写服务，从节点进行数据同步备份，当主节点出现故障后，从节点继续提供服务。

[hash 相关](https://blog.csdn.net/swl1993831/article/details/108023473)


