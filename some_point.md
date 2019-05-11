文件key.go

    internalKey=ukey+7字节的seq+1字节的keyType，seq+keyType构成了一个8字节的uint64类型。
    internalKey使用little-endian字节序存储。
-----------------------------------------
memdb  跳表
http://blog.sina.com.cn/s/blog_72995dcc01017w1t.html
http://jixiuf.github.io/blog/go_package.html/

-- 写数据的时候，是如何保证 wal  存盘的 ？
-- 真正的key 是什么 ？

========
数据保存格式：

如果保存 batch/db.Put  对应一个/几个  chunk
1:journal.Next 表示启动新的chunk
2:chunk格式： 固定7个字节 + batchheader +  sum batch.data
3:batchheader: seq + len( key的个数 ), 共12个。
4：batch.data= add/del + keylen + key + valuelen + value
5:固定7个字节： 4个Byte checksum + 2 Byte len +  1 Byte ChunkType

https://leveldb-handbook.readthedocs.io/zh/latest/basic.html
https://buildmedia.readthedocs.org/media/pdf/leveldb-handbook/latest/leveldb-handbook.pdf

=================
在获取到一个快照之后，leveldb会为本次查询的key构建一个internalKey（格式如上文所述），
其中internalKey的seq字段使用的便是快照对应的seq。
通过这种方式可以过滤掉所有seq大于快照号的数据项。
==== 
如果一个key，在 seq =1，16的时候写入，如果 sep = 10 读取，如何保证读取到 seq = 1 ？
如果有seq = 10的读取，许可 compaction ？

===============

leveldb中使用的cache是一种LRUcache，其结构由两部分内容组成：

    Hash table：用来存储数据；
    LRU：用来维护数据项的新旧信息；

