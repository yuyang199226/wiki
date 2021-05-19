kafka 命令

### 创建topic

```
bin/kafka-topics.sh --zookeeper zookeeper-server:2181 --create --topic my_topic_test --partitions 1 --replication-factor 1
```

会在kafka 的分区日志目录下面新建一个文件夹，这些文件夹的名称由主题名称，破折号（ - ）和分区ID组成因为文件夹名字不能超过255个字符，所以topic 不能超超过249个字符

### 修改topic
增加分区

```
bin/kafka-topics.sh --zookeeper zookeeper-server:2181 --alter --topic my_topic_test --partitions 2
```

增加分区不会更改现有数据的分区。kafka 不会以任何方式重新分配分区数据
目前kafka不支持减少分区

查看topic 列表

```
bin/kafka-topics.sh --zookeeper zookeeper-server:2181 --list
```

发布消息

```
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
```


为什么kafka 吞吐量很大？
1. 顺序访问磁盘
2. 批处理，多个消息合成一组，一起发送。
3. zore-copy。使用linux 的sendfile. 操作系统直接将文件从pagecache 直接发送到网络
4. 压缩一批消息，减少带宽使用。在消费者消费的时候解压

了解数据从文件到套接字的常见数据传输路径就非常重要：

1. 操作系统从磁盘读取数据到内核空间的 pagecache
2. 应用程序读取内核空间的数据到用户空间的缓冲区
3. 应用程序将数据(用户空间的缓冲区)写回内核空间到套接字缓冲区(内核空间)
4. 操作系统将数据从套接字缓冲区(内核空间)复制到通过网络发送的 NIC 缓冲区

### 消息丢失
大多数系统使用消息确认机制
当消息被发送出去，消息被标记为__sent__, 然后 broker 会等待一个来自 consumer 的特定确认，再将消息标记为consumed。这个策略修复了消息丢失的问题，但也产生了新问题。 首先，如果 consumer 处理了消息但在发送确认之前出错了，那么该消息就会被消费两次。第二个是关于性能的，现在 broker 必须为每条消息保存多个状态（首先对其加锁，确保该消息只被发送一次，然后将其永久的标记为 consumed，以便将其移除）。 还有更棘手的问题要处理，比如如何处理已经发送但一直得不到确认的消息。

consumer 
每个consumer 除了共享一个group_id, 每个consumer 都有一个临时且唯一的consumer id
组中的每个consumer用consumer_id注册znode。znode的值包含一个map。这个id只是用来识别在组里目前活跃的consumer，这是个临时节点，如果consumer在处理中挂掉，它就会消失。

consumer offset
consumer 追踪在在每个分区消费的最大的offset。如果offsets.storage=zookeeper 则这个值放到zookeeper 的一个目录下。

replication-factor 表示有几个服务器复制写入的消息。如果您设置了3个复制因子，那么只能最多2个相关的服务器能出问题，否则您将无法访问数据。我们建议您使用2或3个复制因子，以便在不中断数据消费的情况下透明的调整集群


#### 如何通过offset 获取数据?
比如offset = 1008
1. 通过二分法找到小于offset的segment。比如找到00000000000000001000.log和00000000000000001000.index
2. index 文件由两部分组成。相对偏移量：实际物理位置。使用二分法找到偏移量。比如6：45
3. 也就是在偏移量45 对应的1006 ，从log seek(45)开始找，顺序查找知道找到offset=1008的消息。

#### 如何清理过期数据？
1. 删除
2. 压缩

#### kafka 的数据保留策略
1. 按日期时间
2. 按存储大小

#### kafka 可以离开zookeeper单独使用吗？
不能

#### ack 有哪几种？分别代表什么含义？
1. ack=0, 表示producer 不等待broker 返回
2. ack=1, producer 等待 broker 的应答。partition 的 leader罗盘后返回成功，不会等待follower 同步成功
3. ack=-1, producer等待broker的ack，partition的leader和follower全部落盘成功后才返回ack，数据一般不会丢失，延迟时间长但是可靠性高。

生产环境主要是ack=-1 为主，如果压力过大，可以将ack=1,只在测试环境设置 ack=0

消费者组的操作

查看消费者组的offset

```
./bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group console-consumer-30368
```

查看消费者组列表

```
./bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list
```

从控制台接收消息

```
./bin/kafka-console-consumer.sh  --bootstrap-server localhost:9092 --topic store-business-engine-testMetricsUpdateOnline
```

会保存consumer组在每个分区的偏移量
如果consumer 在收消息时没有指定组，那么broker 会自动给他生成一个组。
老版本的位移是提交到zookeeper中的，目录结构是：/consumers/<group.id>/offsets/<topic>/<partitionId>，但是zookeeper其实并不适合进行大批量的读写操作，尤其是写操作。

kafka 什么时候 rebalance 
1. 订阅的主题数量发生变化
2. 主题的分区数量发生变化
3. 消费者组的消费者加入或离开或者crash

在Kafka中实现消费的方式是将日志中的分区划分到每一个消费者实例上，以便在任何时间，每个实例都是分区唯一的消费者。维护消费组中的消费关系由Kafka协议动态处理。如果新的实例加入组，他们将从组中其他成员处接管一些 partition 分区;如果一个实例消失，拥有的分区将被分发到剩余的实例。


kafka 如何保证数据不丢失？


设计抢红包？

布隆过滤器
mysql 隔离级别
go channel 底层实现