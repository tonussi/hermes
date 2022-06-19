package communication

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tonussi/hermes/pkg/proxy"
)

type HTTPCommunicator struct {
	fromAddr string
	toAddr   string
}

func NewHTTPCommunicator(
	fromAddr string,
	toAddr string,
) (*HTTPCommunicator, error) {
	return &HTTPCommunicator{
		fromAddr: fromAddr,
		toAddr:   toAddr,
	}, nil
}

func (comm *HTTPCommunicator) Listen(handle proxy.HandleIncomingMessageFunc) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { comm.requestHandler(w, r, handle) })

	err := http.ListenAndServe(comm.fromAddr, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return err
}

func (comm *HTTPCommunicator) Deliver(data []byte) ([]byte, error) {
	// str1 := string(data[:])
	// fmt.Println("String =", str1)

	var actualRequestRecovered *http.Request

	var err error

	requestAsByffer := bytes.NewReader(data)

	requestRecovered := bufio.NewReader(requestAsByffer)

	actualRequestRecovered, err = http.ReadRequest(requestRecovered)

	if err != nil {
		log.Fatalln(err.Error())
	}

	actualRequestRecovered.Header.Set("Host", "http://"+comm.toAddr+actualRequestRecovered.RequestURI)
	actualRequestRecovered.URL.Host = comm.toAddr
	actualRequestRecovered.Host = comm.toAddr
	actualRequestRecovered.RequestURI = ""
	actualRequestRecovered.URL.Scheme = "http"

	client := &http.Client{}

	res, err := client.Do(actualRequestRecovered)

	if err != nil {
		log.Fatalln(err.Error())
	}

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	// var bufferedResponseHolder = &bytes.Buffer{}

	// err = res.Write(bufferedResponseHolder)

	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }

	return bodyBytes, err
}

func (comm *HTTPCommunicator) requestHandler(w http.ResponseWriter, r *http.Request, handle proxy.HandleIncomingMessageFunc) {
	var bufferedRequestHolder = &bytes.Buffer{}

	var err error

	err = r.Write(bufferedRequestHolder)

	if err != nil {
		log.Fatalln(err.Error())
	}

	resp, err := handle(bufferedRequestHolder.Bytes())

	if err != nil {
		log.Fatalln(err.Error())
	}

	bodyString := string(resp)
	fmt.Fprintf(w, "%+v", bodyString)

	// fmt.Println("String =", bodyString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
