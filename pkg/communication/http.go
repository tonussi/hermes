package communication

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/r3musketeers/hermes/pkg/proxy"
)

type HTTPCommunicator struct {
	fromAddr string
	toAddr   string
	urlPath  string
}

func NewHTTPCommunicator(
	fromAddr string,
	toAddr string,
	connAttempts int,
	connAttemptPeriod time.Duration,
) (*HTTPCommunicator, error) {

	var resp http.Response
	var err error
	for connAttempts > 0 {
		log.Println("connection attempts left:", connAttempts)
		resp, err := http.Get(toAddr)
		if resp.StatusCode != 200 || err != nil {
			connAttempts--
			time.Sleep(connAttemptPeriod)
		}
	}
	if resp.StatusCode != 200 && err != nil {
		return nil, err
	}

	return &HTTPCommunicator{
		fromAddr: fromAddr,
		toAddr:   toAddr,
		urlPath:  "/",
	}, nil
}

func (comm *HTTPCommunicator) Listen(handle proxy.HandleIncomingMessageFunc) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { comm.requestHandler(w, r, handle) })

	err := http.ListenAndServe(comm.fromAddr, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (comm *HTTPCommunicator) Deliver(data []byte) ([]byte, error) {
	// build url to post
	deliveryFullUrlString := comm.addProtocolAndBuildUrl()

	// payload bytes as buffered reader
	bufferedPayload := payloadBytesAsBufferedReader(data)

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
		return nil, err
	}

	// see data that has been returned to the client
	bodyString := string(body)
	log.Print(bodyString)

	return body, err
}

// Extra functions to clean code a little bit

type HttpDelivery struct {
	urlPath string
	payload []byte
}

func (comm *HTTPCommunicator) encapsulateDelivery(r *http.Request, handle proxy.HandleIncomingMessageFunc) HttpDelivery {
	return HttpDelivery{
		urlPath: r.URL.Path,
		payload: comm.getBodyAsBytes(r),
	}
}

func (comm *HTTPCommunicator) getBodyAsBytes(r *http.Request) []byte {
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%b", bodyBytes)

	return bodyBytes
}

func (comm *HTTPCommunicator) requestHandler(w http.ResponseWriter, r *http.Request, handle proxy.HandleIncomingMessageFunc) {
	go comm.handleConnection(r, handle)
}

func (comm *HTTPCommunicator) addProtocolAndBuildUrl() string {
	return "http://" + comm.toAddr + comm.urlPath
}

func payloadBytesAsBufferedReader(data []byte) (ioBufferedValues *bytes.Buffer) {
	return bytes.NewBuffer(data)
}

func (comm *HTTPCommunicator) handleConnection(r *http.Request, handle proxy.HandleIncomingMessageFunc) {
	log.Println("handling connection")

	encapsulateDelivery := comm.encapsulateDelivery(r, handle)

	handle(encapsulateDelivery.payload)
}
