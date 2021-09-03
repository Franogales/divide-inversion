package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		send()
	}
	elapsed := time.Since(start)
	log.Printf("request took %s", elapsed)
}

func send() {
	data := url.Values{
		"name":       {"John Doe"},
		"occupation": {"gardener"},
	}
	resp, err := http.PostForm("http://localhost:8080/credit-asigment", data)
	// resp, err := http.Get("https://ignite3.listentrust.com")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
