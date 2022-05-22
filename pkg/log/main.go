package main

import (
	"log"
	"os"
)

// lib create log file at /tmp
// save batch into log
// get line from log

func main() {
	// Creating an empty file
	// Using Create() function
	myfile, e := os.Create("operations.log")
	if e != nil {
		log.Fatal(e)
	}
	log.Println(myfile)
	myfile.Close()
}
