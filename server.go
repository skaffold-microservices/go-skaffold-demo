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
	fmt.Printf("Starting server at: http://0.0.0.0:%v\n", serverPort)
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":"+serverPort, nil)

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintf(w, "Hello, Friend!")
}

func perpetualLogger() {
	counter := 0
	for {
		counter++
		fmt.Printf("Meaningless log message #%v\n", counter)
		time.Sleep(time.Second * 5)
	}
}
