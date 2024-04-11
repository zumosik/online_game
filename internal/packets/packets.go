package packets

import (
	"errors"

	"github.com/zumosik/bb-marshaling"
)

var ErrInvalidPacketType = errors.New("invalid packet type") // Error for invalid packet type. (If we cant handle it.)

// Packet struct represents a packet of data to be sent over the network.
// It consists of a type and a payload.
// Payloads are defined in the payloads_type.go file.
type Packet struct {
	TypeOfPacket uint8       // TypeOfPacket represents the type of the packet as a single byte.
	Payload      interface{} // Payload can be any type of data.
}

// SerializePacket method takes a Packet and returns it byte representation.
func SerializePacket(p Packet) ([]byte, error) {
	// Serialize payload.
	data := make([]byte, 0)
	var err error

	// To Marshall, we need to get struct type (doesnt work with interface{}).
	switch p.TypeOfPacket {
	case TypeOfPacketConnectReq:
		v := p.Payload.(ConnectReq)
		data, err = bb.Marshall(v)
	case TypeOfPacketConnectResp:
		v := p.Payload.(ConnectResp)
		data, err = bb.Marshall(v)
	case TypeOfPacketNewPlayerConnect:
		v := p.Payload.(NewPlayerConnect)
		data, err = bb.Marshall(v)
	case TypeOfPacketPlayerPosReq:
		v := p.Payload.(PlayerPosReq)
		data, err = bb.Marshall(v)
	case TypeOfPacketPlayerPosResp:
		v := p.Payload.(PlayerPosResp)
		data, err = bb.Marshall(v)
	// add more cases here.
	default:
		return []byte{}, ErrInvalidPacketType
	}

	// Add packet type to start of array.
	packet := append([]byte{p.TypeOfPacket}, data...)

	return packet, err
}

// DeserializePacket method takes a byte array and returns a Packet.
func DeserializePacket(data []byte) (Packet, error) {
	var p Packet

	// Get packet type.
	p.TypeOfPacket = data[0]

	// To Unmarshall we need to get struct type (doesnt work with interface{}).
	var err error
	switch p.TypeOfPacket {
	case TypeOfPacketConnectReq:
		var v ConnectReq
		err = bb.Unmarshall(data[1:], &v)
		p.Payload = v
	case TypeOfPacketConnectResp:
		var v ConnectResp
		err = bb.Unmarshall(data[1:], &v)
		p.Payload = v
	case TypeOfPacketNewPlayerConnect:
		var v NewPlayerConnect
		err = bb.Unmarshall(data[1:], &v)
		p.Payload = v
	case TypeOfPacketPlayerPosReq:
		var v PlayerPosReq
		err = bb.Unmarshall(data[1:], &v)
		p.Payload = v
	case TypeOfPacketPlayerPosResp:
		var v PlayerPosResp
		err = bb.Unmarshall(data[1:], &v)
		p.Payload = v
	// add more cases here.
	default:
		return Packet{}, ErrInvalidPacketType
	}

	return p, err
}
