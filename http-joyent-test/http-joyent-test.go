package main

import (
	"bytes"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {

	// usaage:
	// http-test server
	// http-test client <server-ip:port>

	// get args with this binary stripped off
	args := os.Args[1:]

	if len(args) == 0 {
		panic("No command given in args")
	}

	command := args[0]
	switch command {
	case "server":
		server()
	case "client":
		client()
	default:
		panic("Unknown command: " + command)
	}

}

func server() {

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		<-time.After(time.Second * 2)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func client() {

	// get args with binary and command stripped off
	args := os.Args[2:]

	if len(args) == 0 {
		panic("No server url given in args")
	}

	serverUrl := args[0]

	numGoroutines := 200
	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			endpoint := fmt.Sprintf("%v/%v",
				serverUrl,
				"bar",
			)
			buf := bytes.NewBuffer([]byte(`{"hello":"world"}`))
			resp, err := http.Post(endpoint, "application/json", buf)
			if err != nil {
				log.Printf("Error doing request: %v", err)
				return
			}
			defer resp.Body.Close()
			responseBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading response: %v", err)
				return
			}
			log.Printf("response: %v", string(responseBody))

		}()
	}

	log.Printf("Waiting for goroutines")
	wg.Wait()
	log.Printf("/Waiting for goroutines")

}
