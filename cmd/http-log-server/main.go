package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	listenAddr = flag.String("l", ":8001", "Delivery server address")
)

func whichMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Fprintf(w, "Post from client side!")
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func getBodyAsString(r *http.Request) string {
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)

	log.Printf("%s", bodyString)

	return bodyString
}

func hello(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/db":
		bodyString := getBodyAsString(r)
		fmt.Fprintf(w, "Body: %+v\n", bodyString)
		whichMethod(w, r)
	case "/line":
		bodyString := getBodyAsString(r)
		fmt.Fprintf(w, "Body: %+v\n", bodyString)
		whichMethod(w, r)
	default:
		http.Error(w, "This route is not configured to respond yet....", http.StatusNotFound)
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")

	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
