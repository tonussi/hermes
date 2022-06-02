package communication

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tonussi/studygo/pkg/proxy"
)

type HTTPCommunicator struct {
	fromAddr      string
	toAddr        string
	urlPath       string
	method        string
	requestURI    string
	httpTextBytes []byte
	bodyBytes     []byte
	r             *http.Request
}

func NewHTTPCommunicator(
	fromAddr string,
	toAddr string,
	connAttempts int,
	connAttemptPeriod time.Duration,
) (*HTTPCommunicator, error) {

	http.Get("http://" + toAddr + "/pulse")

	return &HTTPCommunicator{
		fromAddr: fromAddr,
		toAddr:   toAddr,
		urlPath:  "/",
	}, nil
}

func (comm *HTTPCommunicator) Listen(handle proxy.HandleIncomingMessageFunc) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { comm.requestHandler(w, r, handle) })

	err := http.ListenAndServe(comm.fromAddr, nil)

	return err
}

func (comm *HTTPCommunicator) Deliver(data []byte) ([]byte, error) {
	if comm.httpTextBytes == nil {
		return nil, nil
	}

	client := &http.Client{}
	var buf bytes.Buffer
	var res *http.Response
	var req *http.Request
	var err error

	bodyIoReader := bytes.NewBuffer(comm.bodyBytes)
	req, _ = http.NewRequest(comm.r.Method, "http://"+comm.toAddr+comm.r.RequestURI, bodyIoReader)
	res, err = client.Do(req)

	if err != nil {
		log.Println(err.Error())
	}

	defer res.Body.Close()
	res.Write(&buf)

	return buf.Bytes(), err
}

func (comm *HTTPCommunicator) requestHandler(w http.ResponseWriter, r *http.Request, handle proxy.HandleIncomingMessageFunc) {
	comm.r = r
	comm.method = r.Method
	comm.urlPath = r.URL.Path
	comm.requestURI = r.RequestURI
	// Save http request body for later
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	comm.bodyBytes = bodyBytes
	var httpTextBytes bytes.Buffer
	r.Write(&httpTextBytes)
	comm.httpTextBytes = httpTextBytes.Bytes()

	// fmt.Println(httpTextBytes.String())
	resp, _ := handle(comm.httpTextBytes)

	bodyResponseFromAppServer := string(resp)
	// log.Println("Resposta do servidor http-log-server (python)")
	// log.Println(bodyResponseFromAppServer)
	fmt.Fprintf(w, "%s", bodyResponseFromAppServer)
}
