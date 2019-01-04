package main

import (
	"fmt"

	"github.com/davyj0nes/contract-testing-example/consumer/client"
)

func main() {
	c := client.New("http://localhost:8081")
	res, err := c.GetUser("steve")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
