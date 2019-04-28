文件key.go

    internalKey=ukey+7字节的seq+1字节的keyType，seq+keyType构成了一个8字节的uint64类型。
    internalKey使用little-endian字节序存储。
-----------------------------------------
memdb  跳表
http://blog.sina.com.cn/s/blog_72995dcc01017w1t.html
http://jixiuf.github.io/blog/go_package.html/

-- 写数据的时候，是如何保证 wal  存盘的 ？
-- 真正的key 是什么 ？

