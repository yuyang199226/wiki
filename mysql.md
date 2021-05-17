1. 为什么使用索引？
   索引可以提高查询速度
2. mysql 的索引有哪几种？
   聚簇索引和二级索引，每个表都有一个聚簇索引，即使没有设置主键，因为数据在磁盘是有序排列的，所以需要聚簇索引。
https://leetcode-cn.com/circle/discuss/N5PqWI/
3. mysql 使用哪种数据结构实现索引？
   b+ tree
4. b tree 和b+ tree 的区别？
    

# 锁
插入意向锁 ： 是个gap lock, 是执行insert 语句时尝试获取的。加入有索引行4，7 。那可以插入5，6 
next-key 锁


## 事务

### 事务的四个特性
ACID
### 原子性
一个事务中的所有操作要么全部成功，要么都失败。如果事务中间出现错误，会回滚
### 一致性
在事务开始之前和事务结束之后，数据的完整性没有被破坏
### 隔离性
数据库允许多个事务同时对数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据不一致的情况。4种隔离级别： 读未提交，读已提交，可重复读，串行化

### 持久性
一旦事务执行成功，已经写入的数据不会丢失，即使发生硬件故障或者系统故障。

### 隔离级别
- 读未提交
- 读已提交
- 可重复读
- 串行化

脏读： 读到未提交的数据
不可重复读： 两次查询的数据不一样
幻读： 在一次查询的结果集中有一行，但在上次查询是没有这行的 [stackoverflow](https://stackoverflow.com/questions/11043712/what-is-the-difference-between-non-repeatable-read-and-phantom-read)

对于可重复读不能阻止幻读 [mysql](https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_repeatable_read)


## binlog
 binlog 记录DDL ,DML 语句
 用途：
 1. 数据恢复
 2. 主从复制

binlog 什么时候刷到磁盘？
1. 服务器停止的时候
2. 执行flush log 命令
3. 日志文件大于设置的文件最大值

配置：sync_binlog
0: 操作系统自己控制
1 ： 每次事务提交后就刷新到磁盘
N； 完成N个事务后
高版本是默认值是1

binlog 的格式
- statement
- row
- mix

## redo log
是用于在故障恢复的时候修正未完成的事务的基于磁盘的数据结构

重要的参数 innodb_flush_log_at_trx_commit 
0: 每秒刷到磁盘
1: 每次事务提交写入并刷新到磁盘
2: 设置为2时，将在每次事务提交后写入日志，并每秒刷新一次到磁盘。尚未刷新日志的事务可能会在崩溃中丢失。
group commit 

## 多列索引

```
select * from mytable where a=3 or b =4;
# 没有用到索引

(1) select * from mytable where a=3 and b=5 and c=4;
# abc 三列都使用索引，而且都有效

(2) select * from mytable where  c=4 and b=6 and a=3;
# mysql没有那么笨，不会因为书写顺序而无法识辨索引。
# where里面的条件顺序在查询之前会被mysql自动优化，效果跟上一句一样。

(3) select * from mytable where a=3 and c=7;
# a 用到索引，sql中没有使用 b列，b列中断，c没有用到索引

(4) select * from mytable where a=3 and b>7 and c=3;
# a 用到索引，b也用到索引，c没有用到。
# 因为 b是范围索引，所以b处断点，复合索引中后序的列即使出现，索引也是无效的。

(5) select * from mytable where b=3 and c=4;
# sql中没有使用a列， 所以b，c 就无法使用到索引

(6) select * from mytable where a>4 and b=7 and c=9;
# a 用到索引， a是范围索引，索引在a处中断， b、c没有使用索引

(7) select * from mytable where a=3 order by b;
# a用到了索引，b在结果排序中也用到了索引的效果。前面说过，a下面任意一段的b是排好序的

(8) select * from mytable where a=3 order by c;
# a 用到了索引，sql中没有使用 b列，索引中断，c处没有使用索引，在 Extra列 可以看到 filesort

(9) select * from mytable where b=3 order by a;
# 此sql中，先b，后a，导致 b=3 索引无效，排序a也索引无效。
```


分页limit 优化
select * from sv_track_20210426 limit 100000,20;
10 -- 0.01s
100 -- 0.01s
1000 -- 0.39
10000 -- 3.59
100000 -- 31.12

优化后语句
select * from sv_track_20210426 where id >= (select id from sv_track_20210426 where camera_type='FID' limit 100001,1) limit 20;

[慢查询](https://www.cnblogs.com/luyucheng/p/6265594.html)

## 索引合并

以下可能用到索引合并
```
SELECT * FROM tbl_name WHERE key1 = 10 OR key2 = 20;

SELECT * FROM tbl_name
  WHERE (key1 = 10 OR key2 = 20) AND non_key = 30;
```


mysql 是怎么做主从同步的？
mysql 一条更新语句的执行过程？
