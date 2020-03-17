练习目标
=======
go信道 channel

信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，通过使用信道，
数据也可以从一端发送，在另一端接收。

var a chan int
a := make(chan int)

data := <- a // 读取信道 a  
a <- data // 写入信道 a

主协程发生了阻塞，等待信道 done 发送的数据

    func hello(done chan bool) {  
        fmt.Println("Hello world goroutine")
        done <- true
    }
    func main() {  
        done := make(chan bool)
        go hello(done)
        <-done
        fmt.Println("main function")
    }