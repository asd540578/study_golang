package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Name    string
	Age     int
	Balance float64
	Member  bool
}

func main() {
	u1 := User{
		Name:    "Tracy",
		Age:     51,
		Balance: 98.43,
		Member:  true,
	}
	fmt.Println(u1)
	// 顯示 u1 個欄位的記憶體大小和 offset
	fmt.Println("Size/offset:")
	fmt.Println("Name   ", unsafe.Sizeof(u1.Name), unsafe.Offsetof(u1.Name))
	fmt.Println("Age    ", unsafe.Sizeof(u1.Age), unsafe.Offsetof(u1.Age))
	fmt.Println("Balance", unsafe.Sizeof(u1.Balance), unsafe.Offsetof(u1.Balance))
	fmt.Println("Member ", unsafe.Sizeof(u1.Member), unsafe.Offsetof(u1.Member))
	// u1的對其大小和總大小
	fmt.Println("u1 align:", unsafe.Alignof(u1))
	fmt.Println("u1 size :", unsafe.Sizeof(u1))

	// 建立指標指向 u1.Balance
	// u1 (u1.Name) 位置 + 16 + 8＝u1.Balance
	balance := (*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&u1)) + unsafe.Sizeof(u1.Name) + unsafe.Sizeof(u1.Age)))
	*balance += 10000

	// 建立指標指向 u1.Meber
	// u1 (u1.Name) 位置 + 32 ＝ u1.Meber
	member := (*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(&u1)) + unsafe.Offsetof(u1.Member)))
	*member = false

	fmt.Println(u1) // 印出修改後的u1
}
