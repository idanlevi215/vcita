package main

import "fmt"
import "net/http"
import "io/ioutil"
import "log"


//requests Handler

func handler(w http.ResponseWriter, r *http.Request) {
	// read the request body 
    
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err.Error())
    }

  	// print the body
    fmt.Fprintf(w, "%s" , body)
}

func main() {

    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))

}
