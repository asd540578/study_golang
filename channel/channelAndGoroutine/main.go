// 建立go檔案的基礎
package main //定義套件名稱為main

import (
	"fmt"
)

func push(from, to int, out chan int, in chan bool) {
	for i := from; i <= to; i++ {
		<-in     //等待請求
		out <- i //傳回一個數字
	}
}

func main() {
	s1 := 0
	out := make(chan int, 100) //用來接收值的通道
	in := make(chan bool, 100) //越來送出請求的通道

	go push(1, 25, out, in)
	go push(26, 50, out, in)
	go push(51, 75, out, in)
	go push(76, 100, out, in)

	for c := 0; c < 100; c++ {
		in <- true //送出一個請求
		i := <-out //接收一個請求
		fmt.Println(i)
		s1 += i
	}

	fmt.Println("Result:", s1)
}
