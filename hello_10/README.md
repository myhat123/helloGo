练习目标
=======
go工作池

参考资料

https://golangbot.com/buffered-channels-worker-pools/  
https://studygolang.com/articles/12512

WaitGroup

WaitGroup 用于等待一批 Go 协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。

创建了 WaitGroup 类型的变量，其初始值为零值。WaitGroup 使用计数器来工作。

要减少计数器，可以调用 WaitGroup 的 Done() 方法。Wait() 方法会阻塞调用它的 Go 协程，  
直到计数器变为 0 后才会停止阻塞。