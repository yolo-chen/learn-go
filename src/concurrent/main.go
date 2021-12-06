package main

import (
	"fmt"
	"learn-go/src/src/concurrent/runner"
	"time"
)

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("Task complete #%d\n", id)
	}
}

func main() {
	r := runner.New(5 * time.Second)
	r.AddTask(createTask(), createTask(), createTask(), createTask(), createTask(), createTask())
	err := r.Start()
	switch err {
	case runner.ErrInterrupt:
		fmt.Println("task interrupt")
	case runner.ErrTimeout:
		fmt.Println("task timeout")
	default:
		fmt.Println("all task finish")
	}
}
