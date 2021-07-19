package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁加锁后要等待所有读锁执行完毕才能获取

func main() {

	lock := &sync.RWMutex{}
	start := time.Now().Unix()
	sleepSecond := time.Duration(3)

	for i := 0; i < 10; i++ {
		go func(j int) {
			lock.RLock()
			time.Sleep(time.Second * sleepSecond)
			lock.RUnlock()
		}(i)
	}
	time.Sleep(time.Second)

	lock.Lock()
	fmt.Printf("等待了【%d】秒，才获得锁\n", time.Now().Unix()-start)
	lock.Unlock()
}
