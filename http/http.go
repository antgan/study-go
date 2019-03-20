package main

import (
	"fmt"
	"github.com/kimiazhu/golib/safego"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		safego.Go(func() {
			panic("OMG!")
		}, func(err interface{}) {
			fmt.Println("HI1: ", err)
		})
	})
	http.ListenAndServe(":8080", nil)
}
