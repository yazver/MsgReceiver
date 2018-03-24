package main

import (
	"log"

	"github.com/yazver/msgreceiver/shared/lan"
)

func main() {
	server, err := lan.NewServer("127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started.")
	c := server.Receive()
	for m := range c {
		if m != nil {
			log.Println(m.Name)
		} else {
			log.Println("Something wrong.")
		}
	}
}
