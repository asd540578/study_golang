// 建立go檔案的基礎
package main //定義套件名稱為main

import (
	"fmt"
)

func main() { //main()函式
	ch := make(chan int, 1)
	defer close(ch)
	ch <- 5
	i := <-ch
	fmt.Println(i)
}
