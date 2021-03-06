package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	go perpetualLogger()

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "3737"
	}
	externalPort := 30303
	fmt.Printf("Starting server at: http://0.0.0.0:%v\n", externalPort)
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":"+serverPort, nil)

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintf(w, "Hello from go microservice!")
}

func perpetualLogger() {
	counter := 0
	time.Sleep(time.Second * 10)
	for j := 0; j <= 20; j++ { // nothing endless is ever good
		counter++
		fmt.Printf("Meaningless log message #%v\n", counter)
		time.Sleep(time.Second * 5)
	}
	fmt.Printf("Last meaningless message. Server should still run, though.\n")
}
