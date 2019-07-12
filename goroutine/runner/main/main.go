package main

import (
	"github.com/cnmac/golearning/goroutine/runner"
	"log"
	"os"
	"time"
)

// timeout 规定了必须在多少秒内处理完成
const timeout = 3001 * time.Millisecond

func main() {
	log.Println("Starting work.")
	// 为本次执行分配超时时间
	r := runner.New(timeout)
	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())
	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
