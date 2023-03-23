package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

const (
	clientCertName = `.\client_cert.pem`
	serverCertName = `.\server_cert.pem`
	serverKeyName  = `.\server.key`
	host           = "localhost"
	port           = "8080"
)

func main() {
	// 取得客戶端的憑證檔
	clientCert, err := os.ReadFile(clientCertName)
	if err != nil {
		log.Fatal(err)
	}
	// 取得系統的憑證存放區(CertPool)
	clientCAs, err := x509.SystemCertPool()
	if err != nil {
		clientCAs = x509.NewCertPool()
	}
	// 將PEM 格式的客戶端憑證字串加入 CertPool
	if ok := clientCAs.AppendCertsFromPEM(clientCert); !ok {
		log.Println("加入客戶端憑證錯誤")
	}
	// TLS 設定
	tlsConfig := &tls.Config{
		ClientCAs:  clientCAs,                      // 將CertPool 放進信任的 CA 列表
		ClientAuth: tls.RequireAndVerifyClientCert, // 驗證模式
	}

	// 建立 http.Server 結構
	server := &http.Server{
		Addr:      host + ":" + port, // 伺服器網址
		Handler:   nil,               // 用預設的
		TLSConfig: tlsConfig,         // TLS
	}

	log.Println("啟動伺服器")
	http.HandleFunc("/", hello)
	// 啟動HTTPS/LTS 伺服器並仔入伺服器憑證/私鑰
	log.Fatal(server.ListenAndServeTLS(serverCertName, serverKeyName))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("收到請求")
	w.Write([]byte("Hello Golang from a secure server"))
}
