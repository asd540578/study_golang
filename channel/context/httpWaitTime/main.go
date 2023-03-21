package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	c, cancel := context.WithTimeout(r.Context(), time.Second*2) //最多等待兩表
	defer cancel()

	time.Sleep(time.Second * 3) // 等待三秒

	select {
	case <-c.Done(): // 請求超時而被取消
		fmt.Println("Server timeout")
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("<h1>Server timeout</h1>"))
		return
	default: //沒超時正常顯示
		fmt.Println("Hello Golang")
		w.Write([]byte("<h1>Hello Golang</h1>"))
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
