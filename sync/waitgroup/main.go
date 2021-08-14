package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//https://www.toutiao.com/a6775096809460072971/?channel=&source=search_tab
//https://www.toutiao.com/a6698610775587553795/?channel=&source=search_tab
// 内存对齐， 对齐系数

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			s := rand.Intn(3)
			time.Sleep(time.Second * time.Duration(s))
			fmt.Printf("任务【%d】\n", index)
		}(i)
	}

	wg.Wait()
	fmt.Printf("所有任务执行完毕\n")
}
