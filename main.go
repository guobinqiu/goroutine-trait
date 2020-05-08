package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

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
//
//G-M-P三者的关系
//M-P 一对一
//P-G 一对多
//https://www.cnblogs.com/secondtonone1/p/11803961.html
//
//概念
//协程：非抢占式轻量级线程
//非抢占式：非抢占式让原来正在运行的进程继续运行，直至该进程完成或发生某种事件（如I/O请求），才主动放弃处理机
func main() {
	//看源码太累，黑盒测试找规律吧
	//oneCoreWithoutIO()    //从输出结果看，单核并发（concurrency），后面应该有一个队列，每个协程从队列里依次取出来执行。让原来正在运行的协程继续运行，直至该协程完成
	//multiCoresWithoutIO() //从输出结果看，多核并行（parellel）已经看不出队列效果
	oneCoreWithIO()         //从输出结果看，发生了io请求，当前协程主动挂起，让其他协程运行
	//multiCoresWithIO()
}

//一核没io
func oneCoreWithoutIO() {
	runtime.GOMAXPROCS(1)
	go withoutIO("g1")
	go withoutIO("g2")
	go withoutIO("g3")
	time.Sleep(5 * time.Second)
}

//多核没io
func multiCoresWithoutIO() {
	go withoutIO("g1")
	go withoutIO("g2")
	go withoutIO("g3")
	time.Sleep(5 * time.Second)
}

//一核有io
func oneCoreWithIO() {
	runtime.GOMAXPROCS(1)
	go withIO("g1")
	go withIO("g2")
	go withIO("g3")
	time.Sleep(5 * time.Second)
}

//多核有io
func multiCoresWithIO() {
	go withIO("g1")
	go withIO("g2")
	go withIO("g3")
	time.Sleep(5 * time.Second)
}

func withoutIO(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
	}
}

func withIO(name string) {
	for i := 0; i < 10; i++ {
		http.Get("http://www.baidu.com")
		fmt.Println(name, ":", i)
	}
}
