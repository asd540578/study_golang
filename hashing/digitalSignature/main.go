package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	data := []byte("My secret ...")
	// 產生公鑰 私鑰(使用 crypto/rand 產生亂數)
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 用私鑰產生數位簽章
	signedData := ed25519.Sign(privateKey, data)

	// 用公鑰、資料和數位簽章來驗證簽章是否有效
	verfied := ed25519.Verify(publicKey, data, signedData)
	fmt.Println("驗證:", verfied)
}
