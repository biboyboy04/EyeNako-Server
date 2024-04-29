package main

import (
	"fmt"
	"log"

	"github.com/biboyboy04/EyeNako-Server/cmd/api"
)

func main() {
	fmt.Println("Hello Go")

	addr := "localhost:5555"
	server := api.NewAPIServer(addr, nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}