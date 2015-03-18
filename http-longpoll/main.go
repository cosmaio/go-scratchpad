package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/longpoll", func(response http.ResponseWriter, r *http.Request) {

		flush := func() {
			if flusher, ok := response.(http.Flusher); ok {
				flusher.Flush()
			}
		}

		response.Header().Set("Content-Type", "application/json")
		response.Write([]byte("{\"results\":[\r\n"))
		flush()

		heartbeatTicker := time.NewTicker(time.Duration(5) * time.Second)
		defer heartbeatTicker.Stop()
		heartbeat := heartbeatTicker.C

		timeoutTicker := time.NewTicker(time.Duration(30) * time.Second)
		defer timeoutTicker.Stop()
		timeout := timeoutTicker.C

	loop:
		for {
			var err error
			select {

			case <-heartbeat:
				_, err = response.Write([]byte("\n"))
				log.Printf("Wrote a newline to the response")
				flush()
			case <-timeout:
				log.Printf("Timed out, breaking out of loop")
				break loop
			}
			if err != nil {
				log.Printf("Write error: %v", err)
				return // error is probably because the client closed the connection
			}
		}

		response.Write([]byte("]\n"))
		log.Printf("Wrote end of response")

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
