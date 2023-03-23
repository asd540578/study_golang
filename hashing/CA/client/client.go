package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	clientCertName = `.\client_cert.pem`
	clientKeyName  = `.\client.key`
	serverCertName = `.\server_cert.pem`
	host           = "localhost"
	port           = "8080"
)

func main() {
	// 載入客戶端憑證以及私鑰，產生城tls.Certificate物件
	cert, err := tls.LoadX509KeyPair(clientCertName, clientKeyName)
	if err != nil {
		log.Fatal(err)
	}
	// 取得伺服器憑證
	serverCert, err := os.ReadFile(serverCertName)
	if err != nil {
		log.Fatal(err)
	}
	// 取得系統憑證存放區或建立一個新的CertPool
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		rootCAs = x509.NewCertPool()
	}
	// 將PEM格式的伺服器憑證加入CertPool
	if ok := rootCAs.AppendCertsFromPEM(serverCert); !ok {
		log.Fatal("加入伺服器憑證錯誤")
	}
	// TLS 設定，放入客戶端憑證以及信任的伺服器CA清單
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      rootCAs,
	}
	// 建立 http.Client 結構，設定其傳輸層參數使用前面的TLS
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// 客戶端送出請求
	resp, err := client.Get("https://" + host + ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 讀取伺服器回應
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("收到回應:", string(data))
}
