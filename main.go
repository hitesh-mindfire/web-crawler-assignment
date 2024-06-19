package main

import (
	"fmt"
	"web-crawler-assignment/parser"
)

func main() {
	data := parser.Parse()
	fmt.Println("web crawler assignment", data)
}
