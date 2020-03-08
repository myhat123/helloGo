参考资料：

https://golang.google.cn/doc/install
https://golangbot.com/hello-world-gomod/
https://studygolang.com/subject/2

tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz

.bashrc增加配置
export PATH=$PATH:/usr/local/go/bin

1)第一种运行方式
--------------
go build hello.go

$ ./hello

2)第二种运行方式
--------------
安装到指定目录

export GOBIN=~/go/bin
export PATH=$PATH:$GOBIN
go install

在 ~/go/bin 中 hello_01 可执行文件
$ ls go/bin
hello_01

$ hello_01

这种方式和go 1.13之前的版本不同，我们在金融网点运营监控系统中导入程序采用的gb构建

3)第三种运行方式
--------------
go build

生成 hello_01 可执行文件，与 go build hello.go 生成 hello 不同

4)第四种方式
-----------
go run hello.go
go run --work hello.go
