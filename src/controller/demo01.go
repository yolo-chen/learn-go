package main

import (
	"fmt"
	. "learn-go/src/src/stru"
	. "learn-go/src/src/structs"
	"strconv"
	"time"
)

type Teacher struct {
	id    int
	name  string
	age   int
	class string
}

type foo struct {
	num int
}

func main() {
	var arr = make([]*foo, 1)
	arr[0] = &foo{0}
	fmt.Printf("%p\n", arr[0])
	var b = [3]foo{foo{1}, foo{2}, foo{3}}
	fmt.Println("************************")

	for i, p := range b { // 在循环里，i和p都是有一个固定内存空间的变量，在循环时 i p 的值会被不断覆盖
		fmt.Printf("p:------%p\n", &p) // 所以可以看到打印的 i p 的地址值都是固定的，但是其存储的值会随着每一次遍历而更新
		fmt.Printf("i:------%p\n", &i)
		arr = append(arr, &b[i])
		fmt.Printf("%p\n", &b[i])
	}
	fmt.Println("************************")

	for i, e := range arr {
		fmt.Printf("%p\n", arr[i])
		fmt.Printf("e:=======%p\n", &e)
		fmt.Printf("i:=======%p\n", &i)
	}

	cat := Animal{
		Name:  "",
		Color: "",
	}
	cat.Color = "red"

	lily := Person{}
	lily.Height = 12
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	s := timeSeries[1] - timeSeries[0]
	if s >= duration {
		return duration * 2
	}
	return duration + s
}

func failure() {
	fmt.Println("start...")
	teas := make([]Teacher, 10)
	fmt.Printf("for wai teas len :%v\n", len(teas))
	for i := 0; i < len(teas); i++ { // 因为在循环中teas的长度是不断扩容的，所以不能用将len(teas)作为循环终止的条件，
		tea := Teacher{ // 如果用变量来作为循环终止的条件的话，需要注意变量在循环里的值不可被修改。
			id:    i,
			name:  "lily" + strconv.Itoa(i),
			age:   i + i,
			class: "i",
		}
		teas = append(teas, tea)
		fmt.Printf("teas len :%v\n", len(teas))
		time.Sleep(time.Second * 3)
	}
	for tea := range teas {
		fmt.Println(tea)
	}
}

func right() {
	fmt.Println("start...")
	teas := make([]Teacher, 10)
	fmt.Printf("for wai teas len :%v\n", len(teas))
	length := len(teas)
	for i := 0; i < length; i++ {
		tea := Teacher{
			id:    i,
			name:  "lily" + strconv.Itoa(i),
			age:   i + i,
			class: "i",
		}
		teas = append(teas, tea)
		fmt.Printf("teas len :%v\n", len(teas))
	}
	fmt.Println(teas)
}
