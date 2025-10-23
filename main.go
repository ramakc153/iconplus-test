package main

import (
	"fmt"
	"iconplus-test/practicaltest"
)

func main() {
	fmt.Println("iconplus-test server is running")
	practicaltest.PrintData()
	practicaltest.Response()
	practicaltest.Server()
}
