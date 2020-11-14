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

go mod init hello_15

go build -o bin/hello cmd/hello.go

clickhouse-client -u hzg --password 1234