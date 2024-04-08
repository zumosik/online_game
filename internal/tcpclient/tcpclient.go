package tcpclient

import (
	"net"
	"online_game/internal/packets"
)

type User struct {
	Username string
	Pin      uint32
}

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
func (c *TCPClient) Send(p packets.Packet) error {
	return p.Serialize(c.conn)
}

// Receive method reads a packets.Packet from the TCP connection.
func (c *TCPClient) Receive() (packets.Packet, error) {
	return packets.Deserialize(c.conn)
}

func (c *TCPClient) Close() error {
	return c.conn.Close()
}
