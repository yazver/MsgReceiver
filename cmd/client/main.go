package main

import (
	"log"

	"github.com/yazver/msgreceiver/shared/lan"
)

func main() {
	client, err := lan.NewClient("127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	client.Send(&lan.Message{"Hi 1"})
	client.Send(&lan.Message{"Hi 2"})
	client.Send(&lan.Message{"Hi 3"})
	client.Send(&lan.Message{"Hi 4"})
	client.Send(&lan.Message{"Hi 5"})
	client.Send(&lan.Message{"Hi 6"})
	client.Send(&lan.Message{"Hi 7"})
	//client.Close()

	client, err = lan.NewClient("127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	client.Send(&lan.Message{"Hi2 1"})
	client.Send(&lan.Message{"Hi2 2"})
	client.Send(&lan.Message{"Hi2 3"})
	client.Send(&lan.Message{"Hi2 4"})
	client.Send(&lan.Message{"Hi2 5"})
	client.Send(&lan.Message{"Hi2 6"})
	client.Send(&lan.Message{"Hi2 7"})
	client.Close()
}
