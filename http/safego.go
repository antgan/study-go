package main

/*
Package safego wraps "go" call with additional recover feature to keep
your goroutine away from panic.

example:

origin code:

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			panic("OMG!")
		}()
	})
	http.ListenAndServe(":8080", nil)

even the panic occur inside a new goroutine, it will bring the whole server down!

with safego, you can replace with:

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sofego.Go(func() {
			panic("OMG!")
		})
	})
	http.ListenAndServe(":8080", nil)

*/
// Author: ZHU HAIHUA
// Since: 2016-03-22 15:57
import (
	"fmt"
	"github.com/kimiazhu/golib/stack"
	"os"
)

type Handler func(err interface{})

var DefaultHandler = func(err interface{}) {
	fmt.Fprintf(os.Stderr, "recovered: %s\n%s", err, stack.CallStack(3))
}

// Go run the f with a goroutine and keep it away from panic.
// it will use DefaultHandler if argument handler is nil
func Go(f func(), handler ...Handler) {
	handle := DefaultHandler
	switch len(handler) {
	case 1:
		handle = handler[0]
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				handle(r)
			}
		}()

		f()
	}()
}

// GoWithHandler will run function f(args... interface{}) in a go routines.
// And handler the panic if occur with given handler.
func GoWithHandler(f func(args ...interface{}), handler Handler, args ...interface{}) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				handler(r)
			}
		}()

		f(args...)
	}()
}
