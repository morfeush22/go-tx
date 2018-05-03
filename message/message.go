package message

import (
	"github.com/morfeush22/go-tx/crc"
)

const (
	poly = 0xedb88320
	polyInit = 0xffffffff
	polyFinal = 0xffffffff
)
var lookupTable = crc.GenerateCRCLookupTable(poly)

type Message struct {
	data string
	crc  crc.PolyT
}

// ToByte converts message to byte representation
func (m Message) ToByte() []byte {
	crcByte := m.crc.ToByte()
	return append([]byte(m.data), crcByte...)
}

// NewMessage creates new message
func NewMessage(msg string) *Message {
	sum := crc.GenerateCRC([]byte(msg), lookupTable, polyInit, polyFinal)
	return &Message{msg, sum}
}
