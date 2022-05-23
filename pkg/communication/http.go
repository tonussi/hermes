package communication

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/r3musketeers/hermes/pkg/proxy"
)

type HTTPCommunicator struct {
	listener    *int
	deliverConn *int

	connsMux   sync.RWMutex
	connsCount int

	bufferSize     int
	responseBuffer []byte
}

func urlPathHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entrou urlPathHandler")
}

func NewHTTPCommunicator(
	fromAddr string,
	toAddr string,
	connAttempts int,
	connAttemptPeriod time.Duration,
	bufferSize int,
) (*HTTPCommunicator, error) {
	http.HandleFunc("/", urlPathHandler)

	httpPostUrl := fromAddr

	log.Printf("Starting server %s\n", httpPostUrl)

	if err := http.ListenAndServe(fromAddr, nil); err != nil {
		log.Fatal(err)
	}

	deliverAddr, err := net.ResolveTCPAddr("tcp", toAddr)
	if err != nil {
		return nil, err
	}

	var deliverConn *net.TCPConn

	for deliverConn == nil && connAttempts > 0 {
		log.Println("connection attempts left:", connAttempts)
		deliverConn, err = net.DialTCP("tcp", nil, deliverAddr)
		if err != nil {
			connAttempts--
			time.Sleep(connAttemptPeriod)
		}
	}
	if deliverConn == nil && err != nil {
		return nil, err
	}

	listenAddr, err := net.ResolveTCPAddr("tcp", fromAddr)
	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		return nil, err
	}

	return &HTTPCommunicator{
		listener:    listener,
		deliverConn: deliverConn,

		connsMux:   sync.RWMutex{},
		connsCount: 0,

		bufferSize:     bufferSize,
		responseBuffer: make([]byte, bufferSize),
	}, nil
}

func (comm *HTTPCommunicator) Listen(
	handle proxy.HandleIncomingMessageFunc,
) error {
	defer comm.listener.Close()

	for {
		conn, err := comm.listener.AcceptTCP()
		if err != nil {
			return err
		}

		connID := fmt.Sprintf("%d", comm.connsCount)
		comm.connsCount++

		go comm.handleConnection(connID, conn, handle)
	}
}

func (comm *HTTPCommunicator) Deliver(data []byte) ([]byte, error) {
	comm.deliverConn.Write(data)

	n, err := comm.deliverConn.Read(comm.responseBuffer)

	return comm.responseBuffer[:n], err
}

func (comm *HTTPCommunicator) handleConnection(
	connID string,
	clientConn *net.TCPConn,
	handle proxy.HandleIncomingMessageFunc,
) {
	log.Println("handling connection", connID)

	// starts reading from the client
	buffer := make([]byte, comm.bufferSize)

	for {
		n, err := clientConn.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("client closed connection", connID)
			} else {
				log.Printf("error for connection %s: %s", connID, err.Error())
			}

			return
		}

		resp, err := handle(buffer[:n])
		if err != nil {
			clientConn.Write([]byte(err.Error()))
			continue
		}

		clientConn.Write(resp)
	}
}
