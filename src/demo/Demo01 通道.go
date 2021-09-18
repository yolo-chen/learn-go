package main

import "fmt"

func main() {

	continer := make([]int, 0)

	for i := 0; i < 99; i++ {
		continer = append(continer, i)
	}
	for i, _ := range continer {
		fmt.Println(continer[i])
	}
	if true {
		return
	}
	chanl := make(chan int, 51) // 带缓存的channel

	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println("子go程写入chanl ： = ", i)
			chanl <- i
		}
		close(chanl) // 使用for rang 遍历通道需要关闭
	}()

	for v := range chanl {
		fmt.Println("读取的数据为 := ", v)
	}

}
