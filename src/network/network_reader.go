package network

import (
	"encoding/binary"
	"math"
)

type NetworkReader struct {
	data     []byte
	pos      int
	hasError bool
}

func NewNetworkReader(data []byte) *NetworkReader {
	return &NetworkReader{
		data:     data,
		pos:      0,
		hasError: false,
	}
}

func (rdr *NetworkReader) RemainingBytes() int {
	return len(rdr.data) - rdr.pos
}

func (rdr *NetworkReader) HasError() bool {
	return rdr.hasError
}

func (rdr *NetworkReader) ReadByte() byte {
	if rdr.pos >= len(rdr.data) {
		rdr.setError()
		return 0
	}
	b := rdr.data[rdr.pos]
	rdr.pos++
	return b
}

func (rdr *NetworkReader) ReadBytes(amount int) []byte {
	if rdr.pos+amount >= len(rdr.data) {
		rdr.setError()
		return nil
	}
	bytes := rdr.data[rdr.pos : rdr.pos+amount]
	rdr.pos += amount
	return bytes
}

func (rdr *NetworkReader) ReadBool() bool {
	if rdr.pos >= len(rdr.data) {
		rdr.setError()
		return false
	}
	b := rdr.data[rdr.pos]
	rdr.pos++
	return b == 1
}

func (rdr *NetworkReader) ReadShort() int16 {
	if rdr.pos+2 > len(rdr.data) {
		rdr.setError()
		return 0
	}
	slice := rdr.data[rdr.pos : rdr.pos+2]
	value := int16(binary.BigEndian.Uint16(slice))
	rdr.pos += 2
	return value
}

func (rdr *NetworkReader) ReadInt() int32 {
	if rdr.pos+4 > len(rdr.data) {
		rdr.setError()
		return 0
	}
	slice := rdr.data[rdr.pos : rdr.pos+4]
	value := int32(binary.BigEndian.Uint32(slice))
	rdr.pos += 4
	return value
}

func (rdr *NetworkReader) ReadFloat() float32 {
	if rdr.pos+4 > len(rdr.data) {
		rdr.setError()
		return 0.0
	}
	slice := rdr.data[rdr.pos : rdr.pos+4]
	value := math.Float32frombits(binary.BigEndian.Uint32(slice))
	rdr.pos += 4
	return value
}

func (rdr *NetworkReader) ReadString() string {

	length := int(rdr.ReadShort())
	if length < 0 || rdr.pos+length > len(rdr.data) {
		rdr.setError()
		return ""
	}

	slice := rdr.data[rdr.pos : rdr.pos+length]
	value := string(slice)
	rdr.pos += length
	return value
}

func (rdr *NetworkReader) ReadString32() string {

	length := int(rdr.ReadInt())
	if length < 0 || rdr.pos+int(length) > len(rdr.data) {
		rdr.setError()
		return ""
	}

	slice := rdr.data[rdr.pos : rdr.pos+length]
	value := string(slice)
	rdr.pos += length
	return value
}

func (rdr *NetworkReader) setError() {
	rdr.hasError = true
	rdr.pos = len(rdr.data)
}
