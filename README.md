# Message receiver
The simple application and package written in Go that demonstrates the ability to send messages from clients to the server via a TCP connection.

[Documentation.](https://godoc.org/github.com/yazver/msgreceiver/shared/lan)

Sending messages

```go
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
```

Receiving messages

```go
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
```