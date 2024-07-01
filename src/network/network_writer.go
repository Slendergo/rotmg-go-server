package network

import "encoding/binary"

type NetworkWriter struct {
	data []byte
	pos  int
}

func NewNetworkWriter(messageId byte) *NetworkWriter {
	wtr := &NetworkWriter{
		data: make([]byte, 128),
		pos:  0,
	}
	wtr.WriteInt(0)
	wtr.WriteByte(messageId)
	return wtr
}

func (wtr *NetworkWriter) WriteByte(value byte) {
	wtr.ensureCapacity(1)
	wtr.data[wtr.pos] = value
	wtr.pos++
}

func (wtr *NetworkWriter) WriteBool(value bool) {
	wtr.ensureCapacity(1)
	if value {
		wtr.data[wtr.pos] = 1
	} else {
		wtr.data[wtr.pos] = 0
	}
	wtr.pos++
}

func (wtr *NetworkWriter) WriteShort(value int16) {
	wtr.ensureCapacity(2)
	binary.BigEndian.PutUint16(wtr.data[wtr.pos:wtr.pos+2], uint16(value))
	wtr.pos += 2
}

func (wtr *NetworkWriter) WriteInt(value int32) {
	wtr.ensureCapacity(4)
	binary.BigEndian.PutUint32(wtr.data[wtr.pos:wtr.pos+4], uint32(value))
	wtr.pos += 4
}

func (wtr *NetworkWriter) WriteString(value string) {
	length := len(value)
	wtr.ensureCapacity(length + 2)
	wtr.WriteShort(int16(length))
	copy(wtr.data[wtr.pos:], []byte(value))
	wtr.pos += length
}

func (wtr *NetworkWriter) WriteString32(value string) {
	length := len(value)
	wtr.ensureCapacity(length + 4)
	wtr.WriteInt(int32(length))
	copy(wtr.data[wtr.pos:], []byte(value))
	wtr.pos += length
}

func (wtr *NetworkWriter) ensureCapacity(needed int) {
	if wtr.pos+needed > len(wtr.data) {
		newCapacity := len(wtr.data) * 2
		for newCapacity < wtr.pos+needed {
			newCapacity *= 2
		}
		newData := make([]byte, newCapacity)
		copy(newData, wtr.data)
		wtr.data = newData
	}
}

func (wtr *NetworkWriter) Buffer() []byte {
	length := uint32(wtr.pos)
	binary.BigEndian.PutUint32(wtr.data[:4], length)
	return wtr.data[:wtr.pos]
}
