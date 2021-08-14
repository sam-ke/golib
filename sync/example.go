package sync
//======Bad case======

import (
	"fmt"
	"sync"
)

func Mutex() {
	myMap := make(map[string]string)
	lock := &sync.RWMutex{}

	for i := 10000; i > 0; i-- {
		go func() {
			lock.Lock()
			mapKey := fmt.Sprintf("key%d", i)
			myMap[mapKey] = "value"
			lock.Unlock()
		}()

		go func() {
			_, _ = myMap["key1"]
		}()
	}
}

func WatiGroup()  {
	wg := sync.WaitGroup{}
	
	wg.Add(3)
	
	go func() {
		defer wg.Done()
		//...
	}()
	
	go wgFunc(wg)
	
	wg.Wait()
}

func wgFunc(w sync.WaitGroup)  {
	defer w.Done()
	//
}
