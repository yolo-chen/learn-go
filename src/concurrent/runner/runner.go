package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("cannot finish tasks within the timeout")
	ErrInterrupt = errors.New("received interrupt from OS")
)

//Runner 给定一些列的task 要求在规定的timeout 内跑完，否则报错
// 如果操作系统给了中断信号，也报错
type Runner struct {
	interrupt chan os.Signal // 信号， Ctrl + C
	complete  chan error
	duration  <-chan time.Time

	task []func(int)
}

func New(t time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		duration:  time.After(t), // time.After 在给定时间之后返回一个channel
		task:      make([]func(int), 0),
	}
}

func (r *Runner) AddTask(task ...func(int)) {
	r.task = append(r.task, task...)
}

func (r *Runner) run() error {
	for id, task := range r.task {
		select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return ErrInterrupt
		default:
			task(id)
		}
	}
	return nil
}

func (r *Runner) Start() error {

	signal.Notify(r.interrupt, os.Interrupt) // signal.Notify 将系统的 interrupt 传入到给定的 channel里 ， 即 os.Interrupt -> r.interrupt

	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.duration:
		return ErrTimeout
	}

}
