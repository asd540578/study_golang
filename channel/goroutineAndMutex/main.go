package main

import (
	"fmt"
	"sync"
)

type Workers struct { //Worker結構
	in, out   chan int   // 輸入和輸出通道
	workerNum int        // 最大 Goroutine
	mtx       sync.Mutex // 互斥鎖
}

// 初始化 Workers結構
func (w *Workers) init(maxWorkers, maxData int) {
	// 建立通道
	w.in, w.out = make(chan int, maxData), make(chan int)
	// 建立互斥鎖
	w.mtx = sync.Mutex{}
	for i := 0; i < maxWorkers; i++ {
		w.mtx.Lock()
		w.workerNum++ // 紀錄
		w.mtx.Unlock()
		go w.readThem() // 啟動 Goroutine
	}
}

// 輸入資料
func (w *Workers) addData(data int) {
	w.in <- data
}

// 讀出資料
func (w *Workers) readThem() {
	sum := 0
	for i := range w.in { // 讀取通道 in 直到它關閉和無值
		sum += i
	}
	w.out <- sum //  將自己部分的家總值傳給通道

	w.mtx.Lock()
	w.workerNum--
	w.mtx.Unlock()
	if w.workerNum <= 0 { // 減少到0關閉通道
		close(w.out)
	}
}

// 取得結果
func (w *Workers) gatherResult() int {
	close(w.in) // 關閉通道in
	total := 0
	for i := range w.out { // 讀取通道out 直到他關閉和無值
		total += i
	}
	return total
}

func main() {
	maxWorkers := 10
	maxData := 100
	workers := Workers{}              // 建立結構
	workers.init(maxWorkers, maxData) // 初始化

	for i := 1; i <= maxData; i++ {
		workers.addData(i) // 新增資料
	}
	res := workers.gatherResult() // 取得結果
	fmt.Println(res)
}
