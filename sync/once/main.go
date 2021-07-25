package main

import (
	"fmt"
	"sync"
)

var dbInstance *DBInstance
var dbInstanceOnce sync.Once

type DBInstance struct {
	Name string
	Addr string
}

func (db *DBInstance) ExecSQL(sql string) {
	fmt.Println(sql)
}

func GetDBInstance(onceDo bool, wg *sync.WaitGroup) *DBInstance {
	defer wg.Done()

	if dbInstance == nil {
		if onceDo {
			dbInstanceOnce.Do(func() {
				fmt.Println("【once】初始化")
				dbInstance = &DBInstance{}
			})
		} else {
			fmt.Println("【非once】初始化")
			dbInstance = &DBInstance{}
		}
	}
	return dbInstance
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go GetDBInstance(true, &wg)
	}

	dbInstance = nil
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go GetDBInstance(false, &wg)
	}

	wg.Wait()
}
