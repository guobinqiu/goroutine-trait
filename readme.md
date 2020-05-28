### 概念

协程：非抢占式轻量级线程

非抢占式：非抢占式让原来正在运行的进程继续运行，直至该进程完成或发生某种事件（如I/O请求），才主动放弃处理机


```
//源码 src/runtime/proc.go
// Goroutine scheduler
// The scheduler's job is to distribute ready-to-run goroutines over worker threads.
//
// The main concepts are:
// G - goroutine.
// M - worker thread, or machine.
// P - processor, a resource that is required to execute Go code.
//     M must have an associated P to execute Go code, however it can be
//     blocked or in a syscall w/o an associated P.
//
//G 代表一个Goroutine
//M 内核级线程
//P Processor处理器，用来管理和执行Goroutine
```

G-M-P三者的关系
- M 多个
- P 多个（通过设置GOMAXPROCS参数）
- P-M 一对一
- P-G 一对多

> https://www.cnblogs.com/secondtonone1/p/11803961.html
