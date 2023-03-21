package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func sum(from, to int, wg *sync.WaitGroup, res *int32) {
	for i := from; i <= to; i++ {
		atomic.AddInt32(res, int32(i)) //使用原子操作
	}
	wg.Done()
}

func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}

	wg.Add(4) //新增4個Goroutine
	go sum(1, 25, wg, &s1)
	go sum(26, 50, wg, &s1)
	go sum(51, 75, wg, &s1)
	go sum(76, 100, wg, &s1)
	wg.Wait() //等待所有 Goroutine結束

	fmt.Println(s1)
}
