package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func busi(ch chan int) {
	gin.Default()
	for c := range ch {
		fmt.Println("go task =  ", c, "goroutine count = ", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(task int, ch chan int) {
	wg.Add(1)
	ch <- task
}

func main() {
	max := math.MaxInt64
	for i := 0; i < max; i++ {
		go func() {
			fmt.Println("go task = ", i, "goroutine count =", runtime.NumGoroutine())
		}()
	}
}

func main1() {

	ch := make(chan int)
	//模拟用户需求业务的数量
	for i := 0; i < 3; i++ {
		go busi(ch)
	}
	task_cnt := math.MaxInt8

	// 模拟执行任务
	for i := 0; i < task_cnt; i++ {
		sendTask(i, ch)
	}

	wg.Wait()
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine count ======", runtime.NumGoroutine())
	}
}
