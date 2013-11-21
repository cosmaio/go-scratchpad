package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"net/http"
	"time"
)

type StringWrapper struct {
	wrapped string
}

func (sw StringWrapper) String() string {
	json, err := json.Marshal(sw.wrapped)
	if err != nil {
		panic("oops")
	}
	return fmt.Sprintf("%s", json)
}

func main() {
	fmt.Println("hello")

	expvarMap := expvar.NewMap("main")
	stringWrapper := &StringWrapper{wrapped: "foo"}
	expvarMap.Set("onlyvalue", stringWrapper)

	go func() {
		time.Sleep(time.Second)
		stringWrapper.wrapped = "bar"
		time.Sleep(5 * time.Second)
		stringWrapper.wrapped = "baz"
		time.Sleep(5 * time.Second)
		stringWrapper.wrapped = "blah"

	}()

	http.ListenAndServe(":8080", nil)

}
