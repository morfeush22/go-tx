package qpsk

import "testing"
import "github.com/stretchr/testify/assert"

func TestLookupTableGenerator(t *testing.T) {
	lookupTable := GenerateLookupTable()

	assert.Equal(t, byte(0xf0), lookupTable[0xaa])
	assert.Equal(t, byte(0x0f), lookupTable[0x55])
}
