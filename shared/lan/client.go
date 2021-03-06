package lan

import (
	"bufio"
	"encoding/gob"
	"errors"
	"net"
	"strings"
	"sync"
)

// Client create connection to server and allow send messeges to them
type Client struct {
	lock       sync.Mutex
	connection net.Conn
}

// Send message to server
func (c *Client) Send(s *Message) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	rw := bufio.NewReadWriter(bufio.NewReader(c.connection), bufio.NewWriter(c.connection))
	enc := gob.NewEncoder(rw)
	err := enc.Encode(s)
	if err != nil {
		return err
	}
	err = rw.Flush()
	if err != nil {
		return err
	}
	cmd, err := rw.ReadString('\n')
	if err != nil {
		return err
	}
	cmd = strings.TrimSpace(cmd)
	if cmd != "OK" {
		return errors.New("Status: " + cmd)
	}
	return nil
}

// Close connection
func (c *Client) Close() error {
	return c.connection.Close()
}

// NewClient create new client.
func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, nil
	}
	return &Client{connection: conn}, nil
}
