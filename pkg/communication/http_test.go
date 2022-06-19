package communication

import (
	"fmt"
	"log"
	"testing"
)

func TesMessageReceived(t *testing.T) {
	communicator, err := NewHTTPCommunicator(
		"localhost:8000",
		"localhost:8001",
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(communicator)
}
