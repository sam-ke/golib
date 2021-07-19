package main

import (
	"fmt"
	"sync/atomic"
)

type Persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type AtomicPerson struct {
	Value atomic.Value
}

func main() {
	atomicP := AtomicPerson{
		Value: atomic.Value{},
	}

	P := Persion{
		Name: "Hello world!",
		Age:  100,
	}

	atomicP.Value.Store(P)
	p := atomicP.Value.Load().(Persion)

	fmt.Printf("(P==p = %v),atomicP.Value:%+v\n", P == p, p)
}
