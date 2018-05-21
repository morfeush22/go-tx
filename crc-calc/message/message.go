package message

import (
	"encoding/json"
	"github.com/morfeush22/go-tx/crc-calc/crc"
)

const (
	poly      = 0xedb88320
	polyInit  = 0xffffffff
	polyFinal = 0xffffffff
)

var lookupTable = crc.GenerateCRCLookupTable(poly)

type Message struct {
	Data string
	CRC  crc.PolyT
}

// ToByte converts message to byte representation
func (m Message) ToByte() []byte {
	crcByte := m.CRC.ToByte()
	return append([]byte(m.Data), crcByte...)
}

// Marshalize converts message to JSON representation
func (m Message) Marshalize() ([]byte, error) {
	return json.Marshal(map[string][]byte{
		"data": m.ToByte(),
	})
}

// NewMessage creates new message
func NewMessage(msg string) *Message {
	sum := crc.GenerateCRC([]byte(msg), lookupTable, polyInit, polyFinal)
	return &Message{msg, sum}
}
