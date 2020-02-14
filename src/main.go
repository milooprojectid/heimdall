package main

import (
	"fmt"

	"rsc.io/quote"
)

func hello() string {
	return quote.Glass()
}

func main() {
	fmt.Println(hello())
}
