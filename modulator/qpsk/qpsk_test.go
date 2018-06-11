package qpsk

import "testing"
import "github.com/stretchr/testify/assert"

func TestLookupTableGenerator(t *testing.T) {
	lookupTable := GenerateLookupTable()

	assert.Equal(t, byte(0xf0), lookupTable[0xaa])
	assert.Equal(t, byte(0x0f), lookupTable[0x55])
}

func TestQPSKModulator(t *testing.T) {
	modulator := Modulator{}
	inSignal := []byte{0xaa, 0x55}
	outSignal := modulator.Modulate(inSignal)

	assert.Equal(t, byte(0xf0), outSignal.InPhase[0])
	assert.Equal(t, byte(0x0f), outSignal.Quadrature[0])
}
