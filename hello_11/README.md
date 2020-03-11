参考资料

https://golangbot.com/buffered-channels-worker-pools/
https://studygolang.com/articles/12512

以下示例是金融网点运营监控系统的导入cassandra代码的核心

以数据关键字为基础，分配每一个唯一的数据给 job
然后每个 job 作为一个 worker 进行并发插入 cassandra

这种并发量，占用系统资源较低，速度快，效果好！

工作池
=====

缓冲信道的重要应用之一就是实现工作池。

一般而言，工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。

使用缓冲信道来实现工作池。我们工作池的任务是计算所输入数字的每一位的和。例如，如果输入 234，结果会是 9
（即 2 + 3 + 4）。向工作池输入的是一列伪随机数。

我们工作池的核心功能如下：

    创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
    将作业添加到该输入型缓冲信道中。
    作业完成后，再将结果写入一个输出型缓冲信道。
    从输出型缓冲信道读取并打印结果。

定义两个 struct: 1) Job , 2) Result

    type Job struct {  
        id       int
        randomno int
    }
    type Result struct {  
        job         Job
        sumofdigits int
    }

两个缓冲信道

    var jobs = make(chan Job, 10)  
    var results = make(chan Result, 10)

工作协程的函数 worker

    func worker(wg *sync.WaitGroup) {  
        for job := range jobs {
            output := Result{job, digits(job.randomno)}
            results <- output
        }
        wg.Done()
    }

创建一个 Go 协程的工作池

    func createWorkerPool(noOfWorkers int) {  
        var wg sync.WaitGroup
        for i := 0; i < noOfWorkers; i++ {
            wg.Add(1)
            go worker(&wg)
        }
        wg.Wait()
        close(results)
    }

把作业分配给工作者

    func allocate(noOfJobs int) {  
        for i := 0; i < noOfJobs; i++ {
            randomno := rand.Intn(999)
            job := Job{i, randomno}
            jobs <- job
        }
        close(jobs)
    }
