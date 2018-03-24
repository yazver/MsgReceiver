package lan

import (
	"bufio"
	"encoding/gob"
	"errors"
	"net"
	"strings"
)

type Client struct {
	connection net.Conn
}

func (c *Client) Send(s *Message) error {

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
