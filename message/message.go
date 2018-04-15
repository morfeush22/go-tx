package message

import (
	"github.com/morfeush22/go-tx/crc"
	"encoding/binary"
)

const (
	poly = 0xedb88320
	polyInit = 0xffffffff
	polyFinal = 0xffffffff
)
var lookupTable = crc.GenerateCRCLookupTable(poly)

type Message struct {
	msg string
	crc crc.PolyT
}

func (m Message) ToByte() []byte {
	buff := make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, uint32(m.crc))
	return append([]byte(m.msg), buff...)
}

func NewMessage(msg string) *Message {
	crc := crc.GenerateCRC([]byte(msg), lookupTable, polyInit, polyFinal)
	return &Message{msg, crc}
}
