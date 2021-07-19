package main

import (
	"fmt"
	"sync"
	"time"
)

func log(taskIndex int, msg string) {
	fmt.Printf("第【%d】个任务 %s\n", taskIndex, msg)
}

func main() {
	//任务数
	taskNum := 10

	locker := sync.Mutex{}
	condition := sync.NewCond(&locker)

	for i := 0; i < taskNum; i++ {
		go func(index int) {
			log(index, "已启动")
			condition.L.Lock()
			log(index, "等待被信号被唤醒...")
			condition.Wait()
			log(index, "被唤醒")
			condition.L.Unlock()
		}(i)
	}

	time.Sleep(3 * time.Second)
	condition.Signal()

	time.Sleep(3 * time.Second)
	condition.Signal()

	time.Sleep(3 * time.Second)
	condition.Broadcast()

	time.Sleep(time.Second)
}
