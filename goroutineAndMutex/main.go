// 建立go檔案的基礎
package main //定義套件名稱為main

import (
	"log"
	"sync"
)

// 累積值,Mutex 使用Lock / Unlock
func sum(from, to int, wg *sync.WaitGroup, mtx *sync.Mutex, res *int32) {
	for i := from; i <= to; i++ {
		mtx.Lock()
		*res += int32(i)
		mtx.Unlock()
	}

	if wg != nil {
		wg.Done()
	}
}

func main() { //main()函式
	s1 := int32(0)
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	wg.Add(4)
	go sum(1, 25, wg, mtx, &s1)
	go sum(26, 50, wg, mtx, &s1)
	go sum(51, 75, wg, mtx, &s1)
	go sum(76, 100, wg, mtx, &s1)
	wg.Wait()

	log.SetFlags(0) //設定log輸出時不帶其他資訊(時間、程式名稱)
	log.Println(s1)
}
