package main

//https://zhuanlan.zhihu.com/p/76812714

import (
	"runtime/debug"
	"sync"
)

var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	debug.SetGCPercent(-1)

	deal()
}

func deal() {
	//从冲中获取一个值
	obj := bytePool.Get().(*[]byte)

	//业务逻辑...

	//用完之后放入池中
	(*obj)[0] = byte(1)
	*obj = (*obj)[:0]
	bytePool.Put(obj)
}
