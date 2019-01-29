package main

import (
	"fmt"
	"go-starter-commandhandler/src/application"
)

func main() {
	fmt.Println("running")
	application.CommandListenerService("First call")
}
