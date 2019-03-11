package main

import (
	"errors"
	"fmt"
)

func tryRecover() {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(fmt.Sprint("I don't know what to do: %v", r))
		}
	}()

	panic(errors.New("This is error"))
	//panic(123)
}

func main() {
	tryRecover()
}
