package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func getHash(input, hashType string) string {
	switch hashType {
	case "MD5":
		return fmt.Sprintf("%x", md5.Sum([]byte(input)))
	case "SHA256":
		return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
	case "SHA512":
		return fmt.Sprintf("%x", sha512.Sum512([]byte(input)))
	case "SHA3_512":
		return fmt.Sprintf("%x", sha3.Sum512([]byte(input)))
	default:
		return fmt.Sprintf("%x", sha512.Sum512([]byte(input)))
	}
}

func main() {
	fmt.Println("MD5      :", getHash("Hello World!", "MD5"))
	fmt.Println("SHA256   :", getHash("Hello World!", "SHA256"))
	fmt.Println("SHA512   :", getHash("Hello World!", "SHA512"))
	fmt.Println("SHA3_512 :", getHash("Hello World!", "SHA3_512"))

	fmt.Println("以下開始是密碼檢查方法學習")
	password := "mysecretpassword"
	fmt.Println("密碼的明碼 :", password)

	//用bcrypt 將密碼轉成雜湊值
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("密碼雜湊直:", string(hash))

	//測試密碼是否一致
	testString := "mysecretpassword"
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(testString))
	if err != nil {
		fmt.Println("密碼不符")
	} else {
		fmt.Println("密碼相符")
	}
}
