package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)


type requestData struct{
	Name string`json:"name"`
	Email string `json:"email"`
}


func main(){
	cert,err:=tls.LoadX509KeyPair("./certificate.crt","./private.key")

	if err!=nil{
		log.Fatal(err)
	}

	caCert,err:=ioutil.ReadFile("./ca.crt")
	if err!=nil{
		log.Fatal(err)
	}
	caCertPool:=x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig:=&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs: caCertPool,
	}
	httpClient:=&http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	hittingURL:="https://razopr.com"
	data:=requestData{
		Name:"razorpay",
		Email: "razorpay@gmail.com",
	}
	payloadData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	payload := strings.NewReader(string(payloadData))

	req, err := http.NewRequest("POST", hittingURL, payload)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SUCCESS: Handshake completed")
	fmt.Println(string(responseBody))
}