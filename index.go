package main

import (
	"fmt"
	"stockhive-server/cmd/server"
)

func main() {
	r := server.NewServer()	
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to run server")
	}
}
