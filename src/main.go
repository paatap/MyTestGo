package main

import (
	"fmt"

	serv "github.com/paatap/MyTestGo/serv"
)

func main() {
	fmt.Println("Service Started")
	serv.SayHello()
}
