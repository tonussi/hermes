package main

import (
	"bufio"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"os"
)

var configuration []byte
var secret []byte

func Response(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello")
}

func Status(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "ok")
}

func ReadConfig() {
	fmt.Println("reading config...")
	config, e := ioutil.ReadFile("/configs/config.json")
	if e != nil {
		fmt.Printf("Error reading config file: %v\n", e)
		os.Exit(1)
	}
	configuration = config
	fmt.Println("config loaded!")

}

func ReadSecret() {
	fmt.Println("reading secret...")
	s, e := ioutil.ReadFile("/secrets/secret.json")
	if e != nil {
		fmt.Printf("Error reading secret file: %v\n", e)
		os.Exit(1)
	}
	secret = s
	fmt.Println("secret loaded!")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func create_loggers() {
	fmt.Println("creating a file...")

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()
	w := bufio.NewWriter(f)

	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}

func main() {
	fmt.Println("starting...")
	ReadConfig()
	ReadSecret()
	router := fasthttprouter.New()
	router.GET("/", Response)
	router.GET("/status", Status)

	create_loggers()

	log.Fatal(fasthttp.ListenAndServe(":5000", router.Handler))
}
