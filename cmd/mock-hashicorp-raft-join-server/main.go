package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	listenAddr = flag.String("l", ":9000", "Listen server address")
)

type HttpDelivery struct {
	urlPath string
	payload []byte
}

func getBodyAsBytes(r *http.Request) []byte {
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%b", bodyBytes)

	return bodyBytes
}

func encapsulateDelivery(r *http.Request) HttpDelivery {
	return HttpDelivery{
		urlPath: r.URL.Path,
		payload: getBodyAsBytes(r),
	}
}

func getBodyAsString(r *http.Request) string {
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%b", bodyBytes)

	// https://stackoverflow.com/a/40673073
	return string(bodyBytes[:])
}

func hello(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hashicorp-raft/join":
		bodyString := getBodyAsString(r)
		fmt.Fprintf(w, "Body: %+v\n", bodyString)

		encapsulateDelivery := encapsulateDelivery(r)
		log.Printf("%v+", encapsulateDelivery)
	default:
		http.Error(w, "This route is not configured to respond yet....", http.StatusNotFound)
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", hello)

	listenHttpPostUrl := *listenAddr

	log.Printf("Starting listening from server %s\n", listenHttpPostUrl)

	err := http.ListenAndServe(listenHttpPostUrl, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
