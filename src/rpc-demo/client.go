package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	// 用rpc连接服务器
	conn, err := rpc.Dial("tcp", ":8800")
	if err != nil {
		fmt.Println("Dial err", err)
		return
	}
	defer conn.Close()

	resp := ""
	err = conn.Call("hei.SayHello", "beijing", &resp)
	if err != nil {
		fmt.Println("Call err", err)
		return
	}
	fmt.Println(resp)

}
