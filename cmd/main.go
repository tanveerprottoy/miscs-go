package main

import (
	"log"
	"txp/miscs/number"
)

func main() {
	var r any
	r = number.Add(2, 8)
	log.Println(r)
	r = number.Add(2.5, 8.5)
	log.Println(r)
	r = number.Multiply(5.7, 1.9)
	log.Println(r)
}
