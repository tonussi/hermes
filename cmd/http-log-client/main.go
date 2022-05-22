package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

type Profile struct {
	Operation string `json:"operation"`
	Name      string `json:"name"`
	City      string `json:"city"`
}

type Batch struct {
	Batch []Profile `json:"batch"`
}

var (
	deliveryAddr = flag.String("d", "localhost:8001", "Delivery server address")
)

func buildPostUrl(baseUrl string, urlPath string) string {
	return "http://" + baseUrl + urlPath
}

func makeRequest(targetRequestUrlPath string, ioBufferedValues *bytes.Buffer) {
	// make requests
	resp, err := http.Post(targetRequestUrlPath, "application/json", ioBufferedValues)
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

	post := Profile{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	log.Printf("Profile::name %s", post.Name)
	log.Printf("Profile::city %s", post.City)
}

func main() {
	// parse arguments
	flag.Parse()

	// prepare delivery addr
	httpPostUrl := *deliveryAddr
	log.Printf("Http target url is... %s\n", httpPostUrl)

	// begin :: made up payload
	profile1 := Profile{Operation: "INSERT", Name: "lucas", City: "fln"}
	profile2 := Profile{Operation: "INSERT", Name: "marina", City: "fln"}

	var profiles []Profile

	profiles = append(profiles, profile1)
	profiles = append(profiles, profile2)

	log.Println(profiles)

	// end :: made up payload
	batches := Batch{Batch: profiles}
	log.Println(batches)

	// prepare request
	urlPath := buildPostUrl(httpPostUrl, "/db")
	bytesRepresentation, err := json.MarshalIndent(batches, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	// make request
	makeRequest(urlPath, bytes.NewBuffer(bytesRepresentation))
}
