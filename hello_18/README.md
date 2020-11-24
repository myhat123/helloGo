使用反射机制
==========

reflection

tasks中allocate，分配时共享一个参数

共享包
=====

在hello_14中，已经分离出comm。本示例，更改为common。避免循环import。

共享接口
=======

为了tasks中，统一动作，所以共享接口writeCH，然后再进行动作分派，分派给action。

go mod init hello_18

go build -o bin/hello cmd/hello.go

clickhouse-client -u hzg --password 1234

ReplacingMergeTree
==================

参考资料: http://www.clickhouse.com.cn/topic/5bfccb4953dd87ca52effca7

ClickHouse 在使用 ReplacingMergeTree 进行去重时要注意：  
1. 只能根据主键去重。
2. 去重的时候只会去重相同分区的数据，跨分区不会去重，即使是使用 OPTIMIZE **** XXX FINAL 也不会跨分区去重。如果一定要使用这个特性去重，只能将需要去重的数据放在同一个分区。

goroutine池
===========

https://github.com/panjf2000/ants/blob/master/README_ZH.md

简化原有tasks.go的处理

采用go-clickhouse
=================

https://github.com/mailru/go-clickhouse

因官方的clickhouse-go，对decimal类型支持还没有完成，所以选择go-clickhouse包来解决decimal类型  

只需要使用转字符串后写入，即可。