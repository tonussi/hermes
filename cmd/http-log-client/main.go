package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	deliveryAddr = flag.String("d", ":8001", "Delivery server address")
)

type Payload struct {
	Operation string `json:"Operation"`
	Name      string `json:"Name"`
	City      string `json:"City"`
}

func main() {
	// parse arguments
	flag.Parse()

	// prepare delivery addr
	httpPostUrl := *deliveryAddr
	log.Printf("Http target url is... %s\n", httpPostUrl)

	// prepare data
	params := url.Values{}
	params.Add("Operation", "INSERT")
	params.Add("Name", "Lucas")
	params.Add("City", "FLN")

	// prepare request
	resp, err := http.PostForm(httpPostUrl+"/db", params)
	resp.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	log.Print(bodyString)

	post := Payload{}
	err = json.Unmarshal(body, &post)

	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	log.Printf("Payload::Post added with Operation %s", post.Operation)
	log.Printf("Payload::Post added with Name %s", post.Name)
	log.Printf("Payload::Post added with City %s", post.City)
}
