package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	for i:=0; i< 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			s := rand.Intn(3)
			time.Sleep(time.Second*time.Duration(s))
			fmt.Printf("任务【%d】\n", index)
		}(i)
	}

	a := new(interface{})
	wg.Wait()
	fmt.Printf("所有任务执行完毕\n")
}
