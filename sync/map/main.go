package main

import (
	"fmt"
	"sync"
)

func main() {
	syncMap := sync.Map{}
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", "value2")
	syncMap.Store("key3", "value3")

	//取当个值
	v, ok := syncMap.Load("key1")
	if ok {
		fmt.Printf("syncMap['key1']:%+v\n", v)
	}

	//遍历所有值，当回调函数返回false时停止遍历
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("syncMap['%s']:%+v\n", key, v)
		return true
	})

	//删除一个元素
	syncMap.Delete("key2")
}
