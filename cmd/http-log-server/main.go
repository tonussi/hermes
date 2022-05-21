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

type HttpDelivery struct {
	urlPath string
	payload []byte
}

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

func hello(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/db":
		bodyString := getBodyAsString(r)
		fmt.Fprintf(w, "Body: %+v\n", bodyString)
		whichMethod(w, r)
		encapsulateDelivery := encapsulateDelivery(r)
		log.Printf("%v+", encapsulateDelivery)
	case "/line":
		bodyString := getBodyAsString(r)
		fmt.Fprintf(w, "Body: %+v\n", bodyString)
		whichMethod(w, r)
		encapsulateDelivery := encapsulateDelivery(r)
		log.Printf("%v+", encapsulateDelivery)
	default:
		http.Error(w, "This route is not configured to respond yet....", http.StatusNotFound)
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", hello)

	httpPostUrl := *listenAddr

	log.Printf("Starting server %s\n", httpPostUrl)

	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
