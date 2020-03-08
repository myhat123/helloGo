参考资料
https://golangbot.com/go-packages/

文件分离代码 utils.go

package main

func sum(a, b int) int {
	return a+b 
}

hello.go 和 utils.go 同属于 package main
同一个包内，直接进行调用函数