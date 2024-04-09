package tcpclient

import (
	"errors"
	"fmt"
	"log"
	"net"
	"online_game/internal/packets"
)

var ErrNoDataRead = errors.New("no data read")

type User struct {
	Username string
	Pin      uint32
}

const maxReadSize = 1024

// TCPClient struct represents a client that can send and receive packets over a TCP connection.
type TCPClient struct {
	conn net.Conn
	User User
}

func New(addr string, user User) (*TCPClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return &TCPClient{}, err
	}

	return &TCPClient{conn: conn, User: user}, nil
}

// Send method takes a packets.Packet and sends it over the TCP connection.
// Send method takes a packets.Packet and sends it over the TCP connection.
func (c *TCPClient) Send(p packets.Packet) error {
	data, err := packets.SerializePacket(p)
	if err != nil {
		log.Printf("Error serializing packet: %v", err)
		return err
	}

	n, err := c.conn.Write(data)
	if err != nil {
		log.Printf("Error sending packet: %v", err)
		return err
	}
	log.Printf("Successfully sent %d bytes", n)

	return nil
}

// Receive method reads a packets.Packet from the TCP connection.
// ErrNoDataRead is returned if no data was read. This is not an error.
func (c *TCPClient) Receive() (packets.Packet, error) {
	data := make([]byte, 1024)
	n, err := c.conn.Read(data)
	if err != nil {
		return packets.Packet{}, err
	}

	if n == 0 {
		return packets.Packet{}, ErrNoDataRead
	} else {
		log.Printf("Received data: %v", fmt.Sprintf("%v", data[:n]))
	}

	log.Printf("Successfully recieved %d bytes", n)

	return packets.DeserializePacket(data[:n])
}

func (c *TCPClient) Close() error {
	return c.conn.Close()
}
