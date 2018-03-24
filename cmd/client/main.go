package main

import (
	"fmt"
	"log"

	"github.com/yazver/msgreceiver/shared/lan"
)

func sendMessages(id int) {
	client, err := lan.NewClient("127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	for i := 0; i < 6; i++ {
		m := &lan.Message{fmt.Sprintf("Client: %d; Message: %d", id, i)}
		if err := client.Send(m); err != nil {
			log.Printf("Error: %s\n", err)
		}
		log.Println(m.Name)
	}
}

func main() {
	go sendMessages(3)
	sendMessages(1)
	sendMessages(2)
}
