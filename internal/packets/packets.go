package packets

import (
	"github.com/zumosik/bb-marshaling"
	"io"
	"log"
)

// Packet struct represents a packet of data to be sent over the network.
// It consists of a type and a payload.
type Packet struct {
	TypeOfPacket uint8       // TypeOfPacket represents the type of the packet as a single byte.
	Payload      interface{} // Payload can be any type of data.
}

// Serialize method takes an io.Writer and writes the TypeOfPacket as a single byte,
// followed by the Payload encoded using the bb-marshaling package.
// If there's an error at any point, it returns the error.
func (p Packet) Serialize(w io.Writer) error {
	// Write the TypeOfPacket to the writer.
	_, err := w.Write([]byte{p.TypeOfPacket})
	if err != nil {
		return err
	}

	// Encode the Payload and write it to the writer.
	if err := bb.NewEncoder(w).Encode(p.Payload); err != nil {
		return err
	}

	return nil
}

// Deserialize method takes an io.Reader and reads the TypeOfPacket and Payload from it,
// using the bb-marshaling package for the payload.
// If there's an error at any point, it returns an empty Packet and the error.
func Deserialize(r io.Reader) (Packet, error) {
	// Buffer to hold the packet type.
	var buf [1]byte
	var packet Packet

	// Read the packet type from the reader.
	_, err := r.Read(buf[:])
	if err != nil {
		return Packet{}, err
	}
	packet.TypeOfPacket = buf[0]

	log.Println("Packet type:", packet.TypeOfPacket)

	// Decode the payload from the reader.
	if err := bb.NewDecoder(r).Decode(&packet.Payload); err != nil {
		return Packet{}, err
	}

	log.Printf("Received packet: TypeOfPacket=%v, Payload=%v\n", packet.TypeOfPacket, packet.Payload)

	return Packet{}, nil
}
