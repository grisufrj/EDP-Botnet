package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/command", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "{\"args\":[\"ls\",\"-lha\"]}")
	})

	http.HandleFunc("/send", func(rw http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		log.Printf("%s\n", b)
		fmt.Fprintf(rw, "200 - OK")
	})

	log.Println("Starting Server")
	err := http.ListenAndServe(":5000", nil)
	log.Fatal(err)
}
