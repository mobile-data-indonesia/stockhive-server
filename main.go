package main

import (
	"fmt"
	"log"
	"os"
	"stockhive-server/cmd/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	// fmt.Println("Server running on port: ", port)
	r := server.NewServer()	
	if err := r.Run(":" + port); err != nil {
		fmt.Println("Failed to run server")
	}
}
