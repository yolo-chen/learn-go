package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {

	var s = []int{5, 6, 1, 7, 3, 4}
	fmt.Printf("%v", s)
	sort.Ints(s)
	for _, v := range s {
		fmt.Print(v)
	}
	if true {
		return
	}
	listen, err := net.Listen("tcp", ":8999")
	if err != nil {
		panic(err)
	}
	conn, err := listen.Accept()
	if err != nil {
		panic(err)
	}
	for {
		b := make([]byte, 1024)

		read, err := conn.Read(b)
		if err != nil {
			panic(err)
		}
		fmt.Printf("message :%s", b[:read])
	}

}
