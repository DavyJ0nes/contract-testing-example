package main

import (
	"fmt"

	"github.com/davyj0nes/contract-testing-example/consumer/client"
)

func main() {
	c := client.New("http://localhost:8081/data")
	res, err := c.Call()
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
