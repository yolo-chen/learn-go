package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func main() {
	for cost := 10; cost < 15; cost++ {
		start := time.Now()
		password, err := bcrypt.GenerateFromPassword([]byte("password"), cost)
		if err != nil {
			fmt.Println("出错了", err)
		}
		fmt.Printf("password:%s\n", password)

		duration := time.Since(start)
		fmt.Printf("间隔时间：%v\n", duration)
	}
}
