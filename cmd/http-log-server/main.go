package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	listenAddr   = flag.String("l", ":8001", "Listen server address")
	deliveryAddr = flag.String("d", ":8022", "Delivery server address")
)

type HttpDelivery struct {
	urlPath string
	payload []byte
}

// func whichMethod(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		log.Printf("Post from client side!")
// 	default:
// 		log.Printf("Sorry, only POST methods are supported.")
// 	}
// }

func buildPostUrl(baseUrl string, urlPath string) string {
	return "http://" + baseUrl + urlPath
}

func getBodyAsBytes(r *http.Request) []byte {
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%b", bodyBytes)

	return bodyBytes
}

func payloadBytesAsBufferedReader(encapsulateDelivery HttpDelivery) (ioBufferedValues *bytes.Buffer) {
	return bytes.NewBuffer(encapsulateDelivery.payload)
}

func encapsulateDelivery(r *http.Request) HttpDelivery {
	return HttpDelivery{
		urlPath: r.URL.Path,
		payload: getBodyAsBytes(r),
	}
}

func makeRequest(encapsulateDelivery HttpDelivery) {
	deliveryTarget := *deliveryAddr

	// build url to post
	deliveryFullUrlString := buildPostUrl(deliveryTarget, encapsulateDelivery.urlPath)

	// payload bytes as buffered reader
	bufferedPayload := payloadBytesAsBufferedReader(encapsulateDelivery)

	// delivery to a server
	resp, err := http.Post(deliveryFullUrlString, "application/json", bufferedPayload)
	if err != nil {
		log.Fatalln(err)
	}

	// close response body
	defer resp.Body.Close()

	// read body response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	// see data that has been returned to the client
	bodyString := string(body)
	log.Print(bodyString)
}

func hello(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/db":
		// bodyString := getBodyAsString(r)
		// fmt.Fprintf(w, "Body: %+v\n", bodyString)
		// whichMethod(w, r)
		encapsulateDelivery := encapsulateDelivery(r)
		makeRequest(encapsulateDelivery)
		log.Printf("%v+", encapsulateDelivery)
	case "/line":
		// bodyString := getBodyAsString(r)
		// fmt.Fprintf(w, "Body: %+v\n", bodyString)
		// whichMethod(w, r)
		encapsulateDelivery := encapsulateDelivery(r)
		makeRequest(encapsulateDelivery)
		log.Printf("%v+", encapsulateDelivery)
	default:
		http.Error(w, "This route is not configured to respond yet....", http.StatusNotFound)
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", hello)

	listenHttpPostUrl := *listenAddr
	deliveryTarget := *deliveryAddr

	log.Printf("Starting listening from server %s\n", listenHttpPostUrl)
	log.Printf("Starting delivering to server %s\n", deliveryTarget)

	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
