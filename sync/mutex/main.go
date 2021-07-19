package main

import (
	"fmt"
	"sync"
)

type Persion struct {
	execNum int
	Age     int
	mu      *sync.Mutex
	ch      chan struct{}
}

func (p *Persion) ConcurrentExec(task func(p *Persion)) {
	//重置
	p.Age = 0

	for i := 0; i < p.execNum; i++ {
		go task(p)
	}

	for i := 0; i < p.execNum; i++ {
		<-p.ch
	}

	fmt.Printf("过了【%d】年,我的年龄为:%d\n", p.execNum, p.Age)
}

func main() {
	sam := Persion{
		execNum: 10000,
		Age:     0,
		mu:      &sync.Mutex{},
		ch:      make(chan struct{}, 10000),
	}

	//with lock
	sam.ConcurrentExec(func(p *Persion) {
		p.mu.Lock()
		p.Age++
		p.mu.Unlock()

		p.ch <- struct{}{}
	})

	//no lock
	sam.ConcurrentExec(func(p *Persion) {
		//p.mu.Lock()
		p.Age++
		//p.mu.Unlock()

		p.ch <- struct{}{}
	})
}
