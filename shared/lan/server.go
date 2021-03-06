package lan

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
	"net"
)

// Server realise message receiver from clients
type Server struct {
	listener net.Listener
}

// Receive messages from clients
func (s *Server) Receive() chan *Message {
	c := make(chan *Message, 1)

	go func() {

		for {
			conn, err := s.listener.Accept()
			if err != nil {
				close(c)
				return
			}
			go func(conn net.Conn) {
				rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
				defer conn.Close()
				for {
					m := &Message{}
					dec := gob.NewDecoder(rw)
					err = dec.Decode(m)
					if err == io.EOF {
						return
					}
					if err != nil {
						log.Fatal("decode error:", err)
					}
					rw.WriteString("OK\n")
					rw.Flush()
					c <- m
				}
			}(conn)
		}

	}()
	return c
}

// Close connection
func (s *Server) Close() error {
	return s.listener.Close()
}

// NewServer create message's server
func NewServer(address string) (*Server, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &Server{listener: ln}, nil
}
