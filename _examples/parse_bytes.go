package main

import (
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

func _main() error {
	gviz := graphviz.New()
	_, err_1 := graphviz.ParseBytes([]byte("graph test { a -- b }")) // no error
	_, err_2 := graphviz.ParseBytes([]byte("graph test { a -- b"))   // error
	_, err_3 := graphviz.ParseBytes([]byte("graph test { a -- b }")) // no error
	_, err_4 := graphviz.ParseBytes([]byte("graph test { a -- }"))   // error
	_, err_5 := graphviz.ParseBytes([]byte("graph test { a -- c }")) // no error
	_, err_6 := graphviz.ParseBytes([]byte("graph test { a - b }"))  // error
	_, err_7 := graphviz.ParseBytes([]byte("graph test { c -- b }")) // no error

	if err_1 != nil {
		fmt.Println(err_1)
		panic("Test 1 of ParseBytes: Failed")
	}

	if err_2 == nil {
		fmt.Println(err_2)
		panic("Test 2 of ParseBytes: Failed")
	}

	if err_3 != nil {
		fmt.Println(err_3)
		panic("Test 3 of ParseBytes: Failed")
	}

	if err_4 == nil {
		fmt.Println(err_4)
		panic("Test 4 of ParseBytes: Failed")
	}

	if err_5 != nil {
		fmt.Println(err_5)
		panic("Test 5 of ParseBytes: Failed")
	}

	if err_6 == nil {
		fmt.Println(err_6)
		panic("Test 6 of ParseBytes: Failed")
	}

	if err_7 != nil {
		fmt.Println(err_7)
		panic("Test 7 of ParseBytes: Failed")
	}

	gviz.Close()
	return nil
}

func main() {
	if err := _main(); err != nil {
		log.Fatal(err)
	}
}
