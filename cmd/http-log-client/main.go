package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	deliveryAddr = flag.String("d", ":8001", "Delivery server address")
)

func main() {
	// parse arguments
	flag.Parse()

	// prepare delivery addr
	httpPostUrl := *deliveryAddr
	fmt.Println("Http json post url", httpPostUrl)

	// prepare data
	var jsonData = []byte(`{
		"Operation": "INSERT",
		"Name": "Lucas",
		"City": "FLN"
	}`)

	// prepare request
	request, error := http.NewRequest("POST", httpPostUrl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")

	// do request
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	// print results
	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response Body:", string(body))
}
