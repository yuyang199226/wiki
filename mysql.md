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


   
