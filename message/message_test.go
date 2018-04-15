package message

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"encoding/binary"
	"github.com/morfeush22/go-tx/crc"
)

func TestMessageCRC(t *testing.T)  {
	var message = NewMessage("hello_world!")
	expectedCRC := message.crc

	b := message.ToByte()

	data := crc.PolyT(binary.LittleEndian.Uint32(b[len(b) - 4:]))

	assert.Equal(t, expectedCRC, data)
}

func TestMessageByteRepresentation(t *testing.T)  {
	var expectedCRC = []byte{0xAF, 0x62, 0x94, 0x19}
	var message = NewMessage("hello_world!")
	b := message.ToByte()
	for i := range expectedCRC {
		assert.Equal(t, expectedCRC[i], b[len(b) - 4 + i])
	}
}
