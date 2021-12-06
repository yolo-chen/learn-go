package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Phone struct {
}

func main() {

	// 注册rpc服务，绑定对象方法
	err := rpc.RegisterName("hei", new(Phone))
	if err != nil {
		fmt.Println("RegisterName err", err)
		return
	}
	// 设置监听
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()

	fmt.Println("开始监听。。。")

	// 建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept err", err)
		return
	}
	defer conn.Close()

	fmt.Println("连接成功。。。。")

	rpc.ServeConn(conn)

}

func (p *Phone) SayHello(msg string, resp *string) error {
	*resp = msg + "hello"
	return nil
}
