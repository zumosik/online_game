package utils

import (
	"bytes"
	"encoding/binary"
	"math"
)

func WriteFloat32(buf *bytes.Buffer, f float32) {
	WriteUint32(buf, math.Float32bits(f))
}

func WriteFloat64(buf *bytes.Buffer, f float64) {
	WriteUint64(buf, math.Float64bits(f))
}

func WriteBool(buf *bytes.Buffer, b bool) {
	if b == true {
		WriteByte(buf, 1)
	} else {
		WriteByte(buf, 0)
	}
}

// UTF-8
func WriteString(buf *bytes.Buffer, s string) {
	l := uint32(len(s))
	WriteUint32(buf, l)

	buf.Write([]byte(s))
}

func WriteUint16(buf *bytes.Buffer, n uint16) {
	binary.Write(buf, binary.BigEndian, n)
}

func WriteUint32(buf *bytes.Buffer, n uint32) {
	binary.Write(buf, binary.BigEndian, n)
}

func WriteUint64(buf *bytes.Buffer, n uint64) {
	binary.Write(buf, binary.BigEndian, n)
}

func WriteByte(buf *bytes.Buffer, n byte) {
	binary.Write(buf, binary.BigEndian, n)
}

func ReadString(buf *bytes.Buffer) (string, error) {
	n, err := ReadUint32(buf)
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

func ReadFloat32(buf *bytes.Buffer) (float32, error) {
	n, err := ReadUint32(buf)
	return math.Float32frombits(n), err
}

func ReadFloat64(buf *bytes.Buffer) (float64, error) {
	n, err := ReadUint64(buf)
	return math.Float64frombits(n), err
}

func ReadUint16(buf *bytes.Buffer) (n uint16, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func ReadUint32(buf *bytes.Buffer) (n uint32, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func ReadUint64(buf *bytes.Buffer) (n uint64, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func ReadByte(buf *bytes.Buffer) (n byte, err error) {
	err = binary.Read(buf, binary.BigEndian, &n)
	return n, err
}

func ReadBool(buf *bytes.Buffer) (bool, error) {
	n, err := buf.ReadByte()
	if n == 1 {
		return true, err
	}
	return false, err

}
