package message

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"encoding/base64"
)

func TestMessageCRCByteRepresentation(t *testing.T) {
	message := NewMessage("hello_world!")

	expectedData := []byte{'h', 'e', 'l', 'l', 'o', '_', 'w', 'o', 'r', 'l', 'd', '!'}

	bytes := message.ToByte()
	dataSlice := bytes[:len(bytes)-4]

	for i := range expectedData {
		assert.Equal(t, expectedData[i], dataSlice[i])
	}
}

func TestMessageDataByteRepresentation(t *testing.T) {
	message := NewMessage("hello_world!")

	expectedCRC := []byte{0xAF, 0x62, 0x94, 0x19}

	bytes := message.ToByte()
	crcSlice := bytes[len(bytes)-4:]

	for i := range expectedCRC {
		assert.Equal(t, expectedCRC[i], crcSlice[i])
	}
}

func TestMarshalize(t *testing.T) {
	message := NewMessage("hello_world!")
	msg, _ := message.Marshalize()

	var data map[string]interface{}

	json.Unmarshal(msg, &data)
	bytes, _ := base64.StdEncoding.DecodeString(data["data"].(string))

	expectedData := []byte{'h', 'e', 'l', 'l', 'o', '_', 'w', 'o', 'r', 'l', 'd', '!'}
	expectedCRC := []byte{0xAF, 0x62, 0x94, 0x19}

	dataSlice := bytes[:len(bytes)-4]
	crcSlice := bytes[len(bytes)-4:]

	for i := range expectedData {
		assert.Equal(t, expectedData[i], dataSlice[i])
	}

	for i := range expectedCRC {
		assert.Equal(t, expectedCRC[i], crcSlice[i])
	}
}
