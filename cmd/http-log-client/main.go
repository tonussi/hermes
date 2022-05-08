import (
	"io/ioutil"
	"log"
	"net/http"
 )

resp, err := http.Get("https://localhost:5000/log/1")
if err != nil {
   log.Fatalln(err)
}
