package server

import (
	"bytes"
	"encoding/binary"
	"math"
)

func writeFloat32(buf *bytes.Buffer, f float32) {
	writeUint32(buf, math.Float32bits(f))
}

func writeFloat64(buf *bytes.Buffer, f float64) {
	writeUint64(buf, math.Float64bits(f))
}

func writeBool(buf *bytes.Buffer, b bool) {
	if b == true {
		writeByte(buf, 1)
	} else {
		writeByte(buf, 0)
	}
}

// UTF-8
func writeString(buf *bytes.Buffer, s string) {
	l := uint32(len(s))
	writeUint32(buf, l)

	buf.Write([]byte(s))
}

func writeUint16(buf *bytes.Buffer, n uint16) {
	binary.Write(buf, binary.BigEndian, n)
}

func writeUint32(buf *bytes.Buffer, n uint32) {
	binary.Write(buf, binary.BigEndian, n)
}

func writeUint64(buf *bytes.Buffer, n uint64) {
	binary.Write(buf, binary.BigEndian, n)
}

func writeByte(buf *bytes.Buffer, n byte) {
	binary.Write(buf, binary.BigEndian, n)
}

func readString(buf *bytes.Buffer) (string, error) {
	n, err := readUint32(buf)
	if err != nil {
		return "", err
	}

	strBytes := make([]byte, n)
	_, err = buf.Read(strBytes)
	if err != nil {
		return "", err
	}

	str := string(strBytes)

	return str, nil
}

func readFloat32(buf *bytes.Buffer) (float32, error) {
	n, err := readUint32(buf)
	return math.Float32frombits(n), err
}

func readFloat64(buf *bytes.Buffer) (float64, error) {
	n, err := readUint64(buf)
	return math.Float64frombits(n), err
}

func readUint16(buf *bytes.Buffer) (n uint16, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func readUint32(buf *bytes.Buffer) (n uint32, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func readUint64(buf *bytes.Buffer) (n uint64, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func readByte(buf *bytes.Buffer) (n byte, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func readBool(buf *bytes.Buffer) (bool, error) {
	n, err := buf.ReadByte()
	if n == 1 {
		return true, err
	}
	return false, err

}

func (s *Server) isUserIDUnique(userID uint16) bool {
	// Iterate over the playerMap and check if the UserID already exists
	for _, player := range s.playerMap {
		if player.UserID == userID {
			return false // UserID is not unique
		}
	}
	return true // UserID is unique
}
