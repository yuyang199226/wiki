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
### 什么是缓存穿透？
    请求的数据缓存里没有，数据库也没有
### 如何解决缓存穿透？
    第一种解决方法： 如何请求的数据没有，返回一个空值，并用redis 将这个空值缓存起来，设置一个较短的过期时间。
    如果每次请求都是不同的key如何处理？
    采用上面的解决方案会导致redis 存了很多空值，由于redis 的内存淘汰策略，比如lru,lfu，会吧真正有用的数据淘汰。可以用布隆过滤器去解决这个问题。


## 缓存击穿
一个热点key失效，大量的访问打到数据库
### 如何解决缓存击穿
1. 如何量不大，不需要做额外处理
2. 分布式锁
3. 热点key 不设置过期时间
## 缓存雪崩
是大面积版的缓存击穿，key在同一时间失效，大量的访问打到数据库
### 如何解决缓存雪崩
引起缓存雪崩的原因有两点：1. redis 宕机，2. 大量key 在同一时间失效，
1. 给ttl 额外加上一个随机的时间
2. redis 高可用
3. 分级缓存


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

## replication

1. redis 使用异步复制
2. 一个master 可以有多个replicas
3. replicas 可以接受其他replicas 的连接
4. redis 的复制不会阻塞master 端


### 三种主要机制
1. 当master 和 replicas 连接正常时，master 会往replicas 发送一串命令去更新replicas
2. 当连接失败，replicas 会尝试重连并尝试增量更新
3. 如果增量更新失败，就会用全量更新

### 全量更新
1. master fork 出一个子进程去产生一个rdb 文件，同事master 会开启一个buffer,buffer 接收所有的write command。
2. 当rdb 文件生成后，master 吧文件传到relicas, relicas 将其保存到硬盘，并load 到内存中。
3. 然后 master 将 buffer 的command 发给 relicas

#### replication 对于 过期key 的处理
1. relicas 不会主动删除过期key, master 会吧过期key 转成 del 命令发给  relicas
2. relicas 在处理读操作可以识别出过期key,依赖时钟 

### replicas read-only mode
从redis 2.6 开始，默认是只读，对于relicas 的写操作会报错


## redis底层数据结构(redis2.2)
### sds 简单动态字符串

```
struct sdshdr {
    // 等于sds 保存的字符串长度
    unsigned int len;
    // 记录buf 数组未使用的字节的数量
    unsigned int free;
    // 字节数组，用于保存字符串
    char buf[];
};
```

#### sds 与c 字符串的区别
1. O(1) 时间内获取字符串长度
2. 杜绝缓冲区溢出
3. 减少修改字符串时带来的内存重分配次数
4. 二进制安全 ， sds 用len 的属性值判断字符串是不是结束

空间预分配
如果对sds 修改并且需要空间扩展，吐过sds 的长度小于1M， 那sds 的len 的值等于free 的值。大于等于1M就会分配1M的未使用空间
惰性空间释放
当字符串缩短时不立即释放而是用free 记录下来

### 链表

```
typedef struct listNode {
    struct listNode *prev;
    struct listNode *next;
    void *value;
} listNode;
```

```
typedef struct list {
    listNode *head;
    listNode *tail;
    void *(*dup)(void *ptr);
    void (*free)(void *ptr);
    int (*match)(void *ptr, void *key);
    unsigned int len;
} list;
```


