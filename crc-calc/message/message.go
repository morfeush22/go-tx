package message

import (
	"encoding/base64"
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

// Marshal converts message to JSON representation
func (m Message) Marshal() ([]byte, error) {
	return json.Marshal(map[string][]byte{
		"data": m.ToByte(),
	})
}

// Unmarshal converts JSON to byte representation
func (m Message) Unmarshal(message []byte) ([]byte, error) {
	var data map[string]interface{}

	err := json.Unmarshal(message, &data)
	if err != nil {
		return []byte{}, err
	}

	bytes, err := base64.StdEncoding.DecodeString(data["data"].(string))
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

// NewMessage creates new message
func NewMessage(msg string) *Message {
	sum := crc.GenerateCRC([]byte(msg), lookupTable, polyInit, polyFinal)
	return &Message{msg, sum}
}
