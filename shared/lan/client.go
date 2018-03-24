package lan

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type Client struct {
	connection net.Conn
}

func (c *Client) Send(s *Message) {

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(s)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Buffer %d %v\n", b.Len(), b.Bytes())
	_, err = c.connection.Write(b.Bytes())
	// enc := gob.NewEncoder(c.connection)
	// err := enc.Encode(s)
	if err != nil {
		log.Println(err)
	}
	return // err
}

func (c *Client) Close() error {
	return c.connection.Close()
}

func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, nil
	}
	return &Client{connection: conn}, nil
}
