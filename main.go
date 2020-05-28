package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1) //确保单核，这样容易观察

	WithoutIO() //从输出结果看，后面应该有一个队列，每个协程从队列里依次取出来执行。让原来正在运行的协程继续运行，直至该协程完成
	WithIO()    //从输出结果看，后面应该有一个队列，每个协程从队列里依次取出来执行。发生了io请求，当前协程主动挂起，让其他协程运行
	Sched()     //从输出结果看，runtime.Gosched打破了非抢占式特性，手动控制当前协程挂起，让其他协程运行
}

func WithoutIO() {
	go withoutIO("g1")
	go withoutIO("g2")
	go withoutIO("g3")
	time.Sleep(5 * time.Second)
}

func WithIO() {
	go withIO("g1")
	go withIO("g2")
	go withIO("g3")
	time.Sleep(5 * time.Second)
}

func Sched() {
	go sched("g1")
	go sched("g2")
	go sched("g3")
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

func sched(name string) {
	for i := 0; i < 10; i++ {
		if i == 5 {
			runtime.Gosched()
		}
		fmt.Println(name, ":", i)
	}
}
