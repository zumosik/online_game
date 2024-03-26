package server

type Packet struct {
	TypeOfPacket uint8
	Payload      Payload
}

func (p Packet) Serialize() ([]byte, error) {
	b := p.Payload.Serialize()

	res := []byte{p.TypeOfPacket}
	return append(res, b...), nil

	// resp := []byte{p.TypeOfPacket}

	// payloadBytes, err := SerializePayload(p.Payload, p.TypeOfPacket)
	// if err != nil {
	// 	return nil, err
	// }

	// size := len(payloadBytes) + 1 + 4 // 1 is for type, and 4 is for size of packet

	// resp = append(resp, Uint32ToBytes(uint32(size))...)
	// resp = append(resp, payloadBytes...)

	// return resp, err
}

func Deserialize(b []byte) (Packet, error) {
	payload := getPayloadType(b[0])

	if err := payload.Deserialize(b[1:] /* first byte is type of packet */); err != nil {
		return Packet{}, err
	}

	return Packet{
		TypeOfPacket: b[0],
		Payload:      payload,
	}, nil
}

func getPayloadType(packetType byte) Payload {
	var payload Payload

	switch packetType {
	case TypeOfPacketConnectReq:
		payload = &ConnectReq{}
	case TypeOfPacketConnectResp:
		payload = &ConnectResp{}
	case TypeOfPacketPlayerPosReq:
		payload = &PlayerPosReq{}
	case TypeOfPacketPlayerPosResp:
		payload = &PlayerPosResp{}
	}

	return payload
}
