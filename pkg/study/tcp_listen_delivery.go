package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	listenAddr   = flag.String("l", ":8001", "Listen server address")
	deliveryAddr = flag.String("d", ":8022", "Delivery server address")
)

func handleIncomingRequest(listenConn net.Conn, deliveryConn net.Conn) {
	listen_from(listenConn)
	delivery_to(deliveryConn)
}

func listen_from(conn net.Conn) {
	// store incoming data
	buffer := make([]byte, 1024)
	data, err := conn.Read(buffer)
	log.Println(data)
	if err != nil {
		log.Fatal(err)
	}

	// respond
	time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
	conn.Write(buffer)
	conn.Write([]byte("\n"))
	conn.Write([]byte(time))
	conn.Write([]byte("\n"))

	// close conn
	conn.Close()
}

func delivery_to(conn net.Conn) {
	url := "http://" + *deliveryAddr + "/db"
	method := "POST"

	payload := strings.NewReader(`{"batch":[{"operation":"INSERT","name":"Victor Schultz","city":"Domoac"}]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	flag.Parse()
	dialUrl := *deliveryAddr

	listen, errListen := net.Listen("tcp", *listenAddr)
	deliveryConn, errDelivery := net.Dial("tcp", dialUrl)
	if errListen != nil {
		log.Fatal(errListen)
		os.Exit(1)
	}
	if errDelivery != nil {
		log.Fatal(errDelivery)
		os.Exit(1)
	}

	// close listener
	defer listen.Close()
	// close listener
	defer deliveryConn.Close()

	for {
		listenConn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleIncomingRequest(listenConn, deliveryConn)
	}
}
