简单回顾
=======
金融网点运营监控系统

首先要解决的问题，就是将原有在oracle数据库中的数据导入cassandra

2018年2～3月，测试了python进程并发处理，系统占用资源偏大，速度慢
故采用了golang 1.9.6，go协程的方式进行并发处理，效果很好

同时还需要具备两个条件：
1）使用 go-oci8 从oracle读取数据，速度快 https://github.com/mattn/go-oci8
2）写入采用gocql第三方包 https://github.com/gocql/gocql

go-oci8 当时版本，速度还可以，主要是无须过多的类型转换，但现在的版本需要手动控制类型转换，速度降下来了
如果升级到go 1.14版，采用新的mod方式，可以切换到 godror

https://github.com/godror/godror

练习目标
=======
hello_01  go环境及第一个示例
hello_02  函数分离代码
hello_03  文件分离代码
hello_04  go模块机制及包机制分离代码
hello_05  go自带单元测试
hello_06  第三方包集成
hello_07  独立分离主控代码
hello_08  go协程并发初步
hello_09  go信道
hello_10  go工作池
hello_11  go并发作业模式
