package server

import (
	"github.com/zumosik/bb-marshaling"
	"io"
)

type Packet struct {
	TypeOfPacket uint8
	Payload      interface{}
}

// Serialize writes to io.Writer type of packet as 1 byte and then payload using bb-marshaling
func (p Packet) Serialize(w io.Writer) error {
	_, err := w.Write([]byte{p.TypeOfPacket})
	if err != nil {
		return err
	}
	if err := bb.NewEncoder(w).Encode(p.Payload); err != nil {
		return err
	}

	return nil
}

// Deserialize reads from io.Reader type of packet and payload using bb-marshaling
func Deserialize(r io.Reader) (Packet, error) {
	// get packet type
	var buf [1]byte
	var packet Packet
	_, err := r.Read(buf[:])
	if err != nil {
		return Packet{}, err
	}
	packet.TypeOfPacket = buf[0]

	// get payload
	if err := bb.NewDecoder(r).Decode(&packet.Payload); err != nil {
		return Packet{}, err
	}

	return Packet{}, nil
}
