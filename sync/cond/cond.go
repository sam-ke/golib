package main

import (
	"fmt"
	"sync"
	"time"
)

func log(taskIndex int, msg string) {
	task := ""
	if taskIndex >= 0 {
		task = fmt.Sprintf("第【%d】个任务,", taskIndex)
	}
	fmt.Printf("%v, %s %s\n", time.Now().Format("2006-01-02 15:04:05"), task, msg)
}

func main() {
	//任务数
	taskNum := 10

	locker := sync.Mutex{}
	condition := sync.NewCond(&locker)

	for i := 0; i < taskNum; i++ {
		go func(index int) {
			condition.L.Lock()
			log(index, "已启动 等待被信号被唤醒...")
			condition.Wait()
			log(index, "被唤醒")
			condition.L.Unlock()
		}(i)
	}

	time.Sleep(3 * time.Second)
	log(-1, "触发 Signal()")
	condition.Signal()

	time.Sleep(3 * time.Second)
	log(-1, "触发 Broadcast()")
	condition.Broadcast()

	time.Sleep(time.Second)
}
