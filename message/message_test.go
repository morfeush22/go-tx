package message

import "testing"
import (
	"github.com/stretchr/testify/assert"
)

func TestMessageCRCByteRepresentation(t *testing.T)  {
	message := NewMessage("hello_world!")

	expectedData := []byte{'h', 'e', 'l', 'l', 'o', '_', 'w', 'o', 'r', 'l', 'd', '!'}

	bytes := message.ToByte()
	dataSlice := bytes[:len(bytes) - 4]

	for i := range expectedData {
		assert.Equal(t, expectedData[i], dataSlice[i])
	}
}

func TestMessageDataByteRepresentation(t *testing.T)  {
	message := NewMessage("hello_world!")

	expectedCRC := []byte{0xAF, 0x62, 0x94, 0x19}

	bytes := message.ToByte()
	crcSlice := bytes[len(bytes) - 4:]

	for i := range expectedCRC {
		assert.Equal(t, expectedCRC[i], crcSlice[i])
	}
}
