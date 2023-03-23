package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	name string
	age  int
}

type Employee struct {
	name string
	age  int
}

func main() {
	a := User{"John", 42}
	fmt.Printf("a = %#v\n", a)
	//把 a 從 User 轉成 Employee
	b := *(*Employee)(unsafe.Pointer(&a))
	fmt.Printf("b = %#v\n", b)
}
