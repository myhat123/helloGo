练习目标
=======
go协程并发初步

参考资料

https://studygolang.com/articles/12341
https://studygolang.com/articles/12342

并发是指立即处理多个任务的能力。  
并行是指同时处理多个任务。

Go 编程语言原生支持并发。Go 使用 Go 协程（Goroutine） 和信道（Channel）来处理并发。

Go 协程是与其他函数或方法一起并发运行的函数或方法。Go 协程可以看作是轻量级线程。与线程相比，
创建一个 Go 协程的成本很小。因此在 Go 应用中，常常会看到有数以千计的 Go 协程并发地运行。

我们调用了 time 包里的函数 Sleep，该函数会休眠执行它的 Go 协程。在这里，我们使 Go 主协程休眠了 3 秒

```go
func main() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
```

现在，这两个协程并发地运行。numbers 协程首先休眠 250 微秒，接着打印 1，然后再次休眠，打印 2，
依此类推，一直到打印 5 结束。alphabete 协程同样打印从 a 到 e 的字母，并且每次有 400 微秒的休眠时间。 
Go 主协程启动了 numbers 和 alphabete 两个 Go 协程，休眠了 3000 微秒后终止程序。

进程、线程、协程
==============

https://zhuanlan.zhihu.com/p/111346689

进程：进程是系统进行资源分配的基本单位，有独立的内存空间。

线程：线程是 CPU 调度和分派的基本单位，线程依附于进程存在，每个线程会共享父进程的资源。

协程：协程是一种用户态的轻量级线程，协程的调度完全由用户控制，协程间切换只需要保存任务的上下文，没有内核的开销。

线程上下文切换
------------

由于中断处理，多任务处理，用户态切换等原因会导致 CPU 从一个线程切换到另一个线程，切换过程需要保存当前进程的
状态并恢复另一个进程的状态。

上下文切换的代价是高昂的，因为在核心上交换线程会花费很多时间。上下文切换的延迟取决于不同的因素，大概在在 50 到
 100 纳秒之间。考虑到硬件平均在每个核心上每纳秒执行 12 条指令，那么一次上下文切换可能会花费 600 到 1200 条指
 令的延迟时间。实际上，上下文切换占用了大量程序执行指令的时间。

如果存在跨核上下文切换（Cross-Core Context Switch），可能会导致 CPU 缓存失效（CPU 从缓存访问数据的成本
大约 3 到 40 个时钟周期，从主存访问数据的成本大约 100 到 300 个时钟周期），这种场景的切换成本会更加昂贵。

Golang 为并发而生
----------------

Golang 从 2009 年正式发布以来，依靠其极高运行速度和高效的开发效率，迅速占据市场份额。Golang 从语言级别支持
并发，通过轻量级协程 Goroutine 来实现程序并发运行。

Goroutine 非常轻量，主要体现在以下两个方面：

上下文切换代价小： Goroutine 上下文切换只涉及到三个寄存器（PC / SP / DX）的值修改；而对比线程的上下文切换
则需要涉及模式切换（从用户态切换到内核态）、以及 16 个寄存器、PC、SP…等寄存器的刷新；

内存占用少：线程栈空间通常是 2M，Goroutine 栈空间最小 2K；

Golang 程序中可以轻松支持10w 级别的 Goroutine 运行，而线程数量达到 1k 时，内存占用就已经达到 2G。