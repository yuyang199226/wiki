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
